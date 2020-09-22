package run

import (
	c "context"
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/pb"
	"github.com/arpabet/templateserv/pkg/resources"
	"github.com/arpabet/templateserv/pkg/util"
	"github.com/consensusdb/context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	rt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type serverImpl struct {
	ctx                    context.Context
	grpcServer             *grpc.Server
	startTime              time.Time
	Log                    *zap.Logger 			   `inject`
	NodeService 	   		app.NodeService    	   `inject`
	Storage 				app.Storage            `inject`
	ConfigService 			app.ConfigService      `inject`
}

func (t *serverImpl) Run(masterKey string) error {

	t.Log.Info("Start Server",
		zap.String("COS", app.ClassOfService),
		zap.String("NodeId", t.NodeService.NodeIdHex()),
		zap.String("Version", app.GetAppInfo().Version),
		zap.String("Time", time.Now().String()))

	t.startTime = time.Now()

	r := gin.Default()

	if app.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(t.Log, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(t.Log, true))

	signalChain := make(chan os.Signal, 1)
	signal.Notify(signalChain, os.Interrupt)

	api := r.Group("/api")

	api.GET("/status", func(c *gin.Context) {
		appInfo := app.GetAppInfo()

		c.JSON(200, &gin.H{
			"Version":       appInfo.Version,
			"Build":         appInfo.Build,
			"Started":       t.startTime.String(),
			"Node":          t.NodeService.NodeIdHex(),
			"MasterKeyHash": util.GetKeyHash(masterKey),
		})
	})

	api.POST("/stop", func(c *gin.Context) {
		c.String(200, "OK")
		t.Log.Info("Received /stop signal")
		signalChain <- os.Interrupt
	})

	api.POST("/config", func(c *gin.Context) {
		key := c.PostForm("key")
		if key == "" {
			c.String(http.StatusBadRequest, "key is empty")
			return
		}
		value := c.PostForm("value")
		if err := t.ConfigService.Set(key, value); err != nil {
			c.Error(err)
		} else {
			c.String(200, "OK")
		}

	})

	api.GET("/config/:key", func(c *gin.Context) {
		key := c.Param("key")
		if key == "" {
			c.String(http.StatusBadRequest, "key is empty")
			return
		}
		if value, err := t.ConfigService.Get(key); err != nil {
			c.Error(err)
		} else {
			c.String(200, value)
		}
	})

	tlsConfig, err := util.LoadServerConfig(t.Storage)
	if err != nil {
		t.Log.Error("SSL certificates did not find in database", zap.Error(err))
		return err
	}

	tlsAddress, err := t.ConfigService.GetWithDefault(app.ListenTlsAddress, app.DefaultTlsAddress)
	if err != nil {
		t.Log.Error("Failed to get TLS address to Listen", zap.Error(err))
		return err
	}

	server := &http.Server{
		Addr:      tlsAddress,
		TLSConfig: tlsConfig,
		Handler:   r}

	go func() {
		for _ = range signalChain {
			t.Log.Info("Stopping TLS server")
			server.Close()
		}
	}()

	grpcAddress, err := t.ConfigService.Get(app.ListenGrpcAddress)
	if err != nil {
		t.Log.Error("Failed to get gRPC address from config", zap.Error(err))
		return err
	}

	if grpcAddress != "" {

		t.Log.Info("gRPC Server Start", zap.String("grpcAddress", grpcAddress))

		// start listening for grpc
		listenGrpc, err := net.Listen("tcp4", grpcAddress)
		if err != nil {
			t.Log.Fatal("gRPC Port is busy", zap.String("grpcAddress", grpcAddress), zap.Error(err))
			return err
		}

		// Create new grpc pkg
		t.grpcServer = grpc.NewServer()

		// Register services
		pb.RegisterTemplateServiceServer(t.grpcServer, t)

		go t.grpcServer.Serve(listenGrpc)

		grpcGatewayAddress, err := t.ConfigService.Get(app.ListenGrpcGatewayAddress)
		if err != nil {
			t.Log.Error("Failed to get gRPC gateway from config", zap.Error(err))
			return err
		}

		if grpcGatewayAddress != "" {

			t.Log.Info("gRPC Gateway Server Start", zap.String("grpcGatewayAddress", grpcGatewayAddress))

			gatewayServer, err := NewGrpcGatewayServer(c.Background(), grpcAddress, grpcGatewayAddress)
			if err != nil {
				t.Log.Error("Failed to get initialize gRPC gateway server", zap.Error(err))
				return err
			}
			go gatewayServer.ListenAndServe()

		}

	}

	return server.ListenAndServeTLS("", "")
}

func NewGrpcGatewayServer(ctx c.Context, grpcAddress, grpcGatewayAddress string) (*http.Server, error) {

	mux := http.NewServeMux()

	gw := rt.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterTemplateServiceHandlerFromEndpoint(ctx, gw, grpcAddress, opts)
	if err != nil {
		return nil, err
	}
	mux.Handle("/v1/", gw)
	mux.HandleFunc("/", serveWelcome)
	mux.Handle("/swagger/", http.FileServer(resources.AssetFile()))
	//mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{Addr: grpcGatewayAddress, Handler: mux}, nil

}

var welcomeTpl = util.MustAssetTemplate("templates/welcome.tmpl")

func serveWelcome(w http.ResponseWriter, r *http.Request) {
	welcomeTpl.Execute(w, r)
}
