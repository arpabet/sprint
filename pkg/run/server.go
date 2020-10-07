/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package run

import (
	c "context"
	"github.com/arpabet/context"
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/pb"
	"github.com/arpabet/sprint/pkg/util"
	rt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type serverImpl struct {
	ctx        context.Context
	nodeServer *grpc.Server
	grpcServer *grpc.Server
	httpServer *http.Server

	Log                    *zap.Logger 			  `inject`
	NodeService  	   	   app.NodeService    	  `inject`
	Storage 			   app.Storage            `inject`
	ConfigService 		   app.ConfigService      `inject`
	DatabaseService 	   app.DatabaseService    `inject`

	startTime             time.Time
	signalChain			  chan os.Signal

	closeOnce             sync.Once
}

func NewServerImpl(ctx  context.Context) *serverImpl {
	return &serverImpl{
		ctx: ctx,
		startTime: time.Now(),
		signalChain: make(chan os.Signal, 1),
	}
}

func (t *serverImpl) Run(masterKey string) error {

	t.Log.Info("Start Server",
		zap.String("COS", app.ClassOfService),
		zap.String("NodeId", t.NodeService.NodeIdHex()),
		zap.String("Version", app.Version),
		zap.Time("Time", t.startTime))


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

	return t.nodeServer.Serve(listenNode)
}

func NewHttpServer(ctx c.Context, httpAddress, grpcAddress string) (*http.Server, error) {

	mux := http.NewServeMux()

	v1 := rt.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if grpcAddress != "" {
		var err error
		if app.RegisterGatewayServices != nil {
			err = app.RegisterGatewayServices(ctx, v1, grpcAddress)
		} else {
			err = pb.RegisterExampleServiceHandlerFromEndpoint(ctx, v1, grpcAddress, opts)
		}
		if err != nil {
			return nil, err
		}
	}
	mux.Handle("/v1/", v1)
	mux.HandleFunc("/", serveWelcome)
	mux.Handle("/swagger/", http.FileServer(app.Resources))

	if app.Endpoints != nil {
		for _, entry := range app.Endpoints {
			mux.Handle(entry.Pattern, entry.Handler)
		}
	}

	//mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{Addr: httpAddress, Handler: mux}, nil

}

var welcomeTpl = util.MustAssetTemplate("templates/welcome.tmpl")

func serveWelcome(w http.ResponseWriter, r *http.Request) {
	welcomeTpl.Execute(w, r)
}

func (t *serverImpl) Close() {
	t.closeOnce.Do(func() {
		t.Log.Info("Server Stop", zap.Time("End", time.Now()))
		if t.httpServer != nil {
			t.httpServer.Close()
		}
		if t.grpcServer != nil {
			t.grpcServer.Stop()
		}
		t.nodeServer.Stop()
	})
}