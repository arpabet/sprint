/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package run

import (
	c "context"
	"fmt"
	"github.com/arpabet/context"
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/pb"
	"github.com/arpabet/sprint/pkg/util"
	rt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
	"github.com/fsnotify/fsnotify"
)

type serverImpl struct {
	ctx        			   context.Context
	nodeServer 			   *grpc.Server
	grpcServer 			   *grpc.Server
	httpServer 			   *http.Server

	Log                    *zap.Logger 			  `inject`
	NodeService  	   	   app.NodeService    	  `inject`
	Storage 			   app.Storage            `inject`
	ConfigService 		   app.ConfigService      `inject`
	DatabaseService 	   app.DatabaseService    `inject`

	startTime             time.Time
	signalChain			  chan os.Signal

	closeOnce             sync.Once
	restarting            atomic.Bool

	autoupdate            *fsnotify.Watcher
	autoupdateDone        chan bool
	distrStat             os.FileInfo
	requestUpdateTimestamp        atomic.Int64

}

func NewServerImpl(ctx  context.Context) *serverImpl {
	srv := &serverImpl{
		ctx: ctx,
		startTime: time.Now(),
		signalChain: make(chan os.Signal, 1),
		autoupdateDone: make(chan bool),
	}
	srv.restarting.Store(false)
	return srv
}

func (t *serverImpl) Run(masterKey string) error {

	t.Log.Info("Start Server",
		zap.String("COS", app.ClassOfService),
		zap.String("NodeId", t.NodeService.NodeIdHex()),
		zap.String("Version", app.Version),
		zap.Time("Time", t.startTime))

	autoupdate, err := t.ConfigService.GetBool(app.Autoupdate)
	if err != nil {
		t.Log.Error("Autoupdate", zap.Error(err))
		return err
	}

	nodeAddress, err := t.ConfigService.GetWithDefault(app.ListenNodeAddress, app.GetNodeAddress())
	if err != nil {
		t.Log.Error("Node Address", zap.String("nodeAddress", nodeAddress), zap.Error(err))
		return err
	}

	t.Log.Info("Node Server",
		zap.Time("Start", time.Now()),
		zap.String("nodeAddress", nodeAddress))

	// start listening for grpc control port
	listenNode, err := net.Listen("tcp4", nodeAddress)
	if err != nil {
		t.Log.Error("Bind Port", zap.String("nodeAddress", nodeAddress), zap.Error(err))
		return err
	}

	nodeTlsConfig, err := util.LoadServerConfig(t.Storage)
	if err != nil {
		t.Log.Error("Node TLS", zap.Error(err))
		return err
	}

	// Create new node server
	t.nodeServer = grpc.NewServer(grpc.Creds(credentials.NewTLS(nodeTlsConfig)))

	// Register control service
	pb.RegisterNodeServiceServer(t.nodeServer, t)

	grpcAddress, err := t.ConfigService.Get(app.ListenGrpcAddress)
	if err != nil {
		t.Log.Error("gRPC Address", zap.String("grpcAddress", grpcAddress), zap.Error(err))
		return err
	}

	if grpcAddress != "" {

		t.Log.Info("gRPC Start", zap.Time("Start", time.Now()), zap.String("grpcAddress", grpcAddress))

		// start listening for grpc port
		listenGrpc, err := net.Listen("tcp4", grpcAddress)
		if err != nil {
			t.Log.Fatal("Bind Port", zap.String("grpcAddress", grpcAddress), zap.Error(err))
			return err
		}

		// Create new grpc pkg
		t.grpcServer = grpc.NewServer()

		// Register services
		if app.RegisterServices != nil {
			if err := app.RegisterServices(t.ctx, t.grpcServer); err != nil {
				t.Log.Error("RegisterServices", zap.String("grpcAddress", grpcAddress), zap.Error(err))
				return err
			}
		} else {
			pb.RegisterExampleServiceServer(t.grpcServer, t)
		}

		go t.grpcServer.Serve(listenGrpc)

	}

	httpAddress, err := t.ConfigService.Get(app.ListenHttpAddress)
	if err != nil {
		t.Log.Error("HTTP Address", zap.String("httpAddress", httpAddress), zap.Error(err))
		return err
	}

	if httpAddress != "" {

		t.httpServer, err = NewHttpServer(c.Background(), httpAddress, grpcAddress)
		if err != nil {
			t.Log.Error("HTTP Server", zap.Error(err))
			return err
		}
		go t.httpServer.ListenAndServe()

	}

	signal.Notify(t.signalChain, os.Interrupt)
	go func() {
		for _ = range t.signalChain {
			t.Close()
		}
	}()

	distrFile := app.GetDistrFile()
	if autoupdate && distrFile != "" {
		err = t.Autoupdate(distrFile)
		if err != nil {
			t.Log.Error("Autoupdate Watcher", zap.String("distrFile", distrFile), zap.Error(err))
		}
	}

	return t.nodeServer.Serve(listenNode)
}

func NewHttpServer(ctx c.Context, httpAddress, grpcAddress string) (*http.Server, error) {

	mux := http.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if grpcAddress != "" {
		api := rt.NewServeMux()
		var err error
		if app.RegisterGatewayServices != nil {
			err = app.RegisterGatewayServices(ctx, api, grpcAddress)
		} else {
			err = pb.RegisterExampleServiceHandlerFromEndpoint(ctx, api, grpcAddress, opts)
		}
		if err != nil {
			return nil, err
		}
		fmt.Printf("Route /api to %v\n", api)
		mux.Handle("/api/", api)
	}

	indexDefined := false
	if app.Endpoints != nil {
		for _, entry := range app.Endpoints {
			if entry.Pattern == "/" {
				indexDefined = true
			}
			if app.IsDev {
				fmt.Printf("Route Entry %s to %v\n", entry.Pattern, entry.Handler)
			}
			mux.Handle(entry.Pattern, entry.Handler)
		}
	}

	//mux.Handle("/metrics", promhttp.Handler())

	assetsFileSys := http.FileServer(app.Assets)
	for _, name := range app.AssetNames {
		pattern := "/" + name
		if pattern == "/" {
			indexDefined = true
		}
		if app.IsDev {
			fmt.Printf("Route Asset %s to app.Assets\n", pattern)
		}
		mux.Handle(pattern, assetsFileSys)
	}

	if !indexDefined {
		index, err := newIndexPage()
		if err != nil {
			return nil, err
		}
		if app.IsDev {
			fmt.Printf("Route Index / to %v\n", index)
		}
		mux.Handle("/", index)
	}

	return &http.Server{Addr: httpAddress, Handler: mux}, nil

}

func (t *serverImpl) Close() {
	t.closeOnce.Do(func() {
		t.Log.Info("Server Stop", zap.Time("End", time.Now()))
		if t.autoupdate != nil {
			t.autoupdateDone <- true
			t.autoupdate.Close()
		}
		if t.httpServer != nil {
			t.httpServer.Close()
		}
		if t.grpcServer != nil {
			t.grpcServer.Stop()
		}
		t.nodeServer.Stop()
	})
}

func (t *serverImpl) Restarting() bool {
	return t.restarting.Load()
}


