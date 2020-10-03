/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package run

import (
	c "context"
	"github.com/arpabet/context"
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/pb"
	"github.com/arpabet/templateserv/pkg/util"
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
	ctx                    context.Context
	controlServer          *grpc.Server
	grpcServer             *grpc.Server
	httpServer  		   *http.Server

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


	controlAddress, err := t.ConfigService.GetWithDefault(app.ListenControlAddress, app.GetControlAddress())
	if err != nil {
		t.Log.Error("Control Address", zap.String("controlAddress", controlAddress), zap.Error(err))
		return err
	}

	t.Log.Info("Control Server",
		zap.Time("Start", time.Now()),
		zap.String("controlAddress", controlAddress))

	// start listening for grpc control port
	listenControl, err := net.Listen("tcp4", controlAddress)
	if err != nil {
		t.Log.Error("Bind Port", zap.String("controlAddress", controlAddress), zap.Error(err))
		return err
	}

	controlTlsConfig, err := util.LoadServerConfig(t.Storage)
	if err != nil {
		t.Log.Error("Control TLS", zap.Error(err))
		return err
	}

	// Create new control server
	t.controlServer = grpc.NewServer(grpc.Creds(credentials.NewTLS(controlTlsConfig)))

	// Register control service
	pb.RegisterControlServiceServer(t.controlServer, t)

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
		pb.RegisterTemplateServiceServer(t.grpcServer, t)
		if app.RegisterServices != nil {
			if err := app.RegisterServices(t.ctx, t.grpcServer); err != nil {
				t.Log.Error("RegisterServices", zap.String("grpcAddress", grpcAddress), zap.Error(err))
				return err
			}
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

	return t.controlServer.Serve(listenControl)
}

func NewHttpServer(ctx c.Context, httpAddress, grpcAddress string) (*http.Server, error) {

	mux := http.NewServeMux()

	gw := rt.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if grpcAddress != "" {
		err := pb.RegisterTemplateServiceHandlerFromEndpoint(ctx, gw, grpcAddress, opts)
		if app.RegisterGatewayServices != nil {
			if err := app.RegisterGatewayServices(ctx, gw, grpcAddress); err != nil {
				return nil, err
			}
		}
		if err != nil {
			return nil, err
		}
	}
	mux.Handle("/v1/", gw)
	mux.HandleFunc("/", serveWelcome)
	mux.Handle("/swagger/", http.FileServer(app.Resources))

	if app.Endpoints != nil {
		for pattern, handler := range app.Endpoints {
			mux.Handle(pattern, handler)
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
		t.httpServer.Close()
		t.grpcServer.Stop()
		t.controlServer.Stop()
	})
}