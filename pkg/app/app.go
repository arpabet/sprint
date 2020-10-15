/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	c "context"
	"encoding/base64"
	"arpabet.pkg.is/context"
	"github.com/arpabet/sprint/pkg/assets"
	"github.com/arpabet/sprint/pkg/resources"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"time"
)

var (

	Version   string
	Build     string

	MasterKey = "SPRINT_MASTER_KEY"
	ClassOfService = os.Getenv("COS")

	ExecutableName = "sprint"
	ApplicationName = "Sprint"
	PackageName = "github.com/arpabet/sprint"

	Copyright = "Copyright (C) 2020-present Arpabet Inc. All rights reserved."

	UserDir = true

	UserDirPerm = os.FileMode(0700)
	DataDirPerm = os.FileMode(0700)
	ExeFilePerm = os.FileMode(0755)

	KeySize = 32  // 256-bit AES key

	NodeIdBits = 48
	NodeIdSize = NodeIdBits / 8

	Encoding = base64.RawURLEncoding

	PlainContentType = "text/plain"

	IsProd = ClassOfService == "prod"
	IsDev = ClassOfService == "dev" || ClassOfService == ""

	StopDelay = time.Millisecond * 10

	KeyRotationDuration = time.Hour * 24 * 7

	DefaultSSLFolder = "ssl"

	EventPrefix = byte('#')
	ConfigPrefix = "config:"
	ConfigPrefixLen = len(ConfigPrefix)

	NodeId = "node.id"

	DefaultNodeAddress = "localhost:6000"
	ListenNodeAddress  = "listen.node.address"
	ListenGrpcAddress  = "listen.grpc.address"        			 // if empty then do not run gRPC server
	ListenHttpAddress  = "listen.http.address"                   // if empty then do not run HTTP server

	Autoupdate  = "autoupdate"
)

type Endpoint struct {
	Pattern   string
	Handler   http.Handler
}

// Hooks
var (

	Scan      []interface{}
	Initialized  func(context.Context) error
	RegisterServices  func(context.Context, *grpc.Server) error
	RegisterGatewayServices  func(ctx c.Context, gw *runtime.ServeMux, grpcAddress string) error

	Endpoints  []Endpoint

	IndexFile = "templates/index.tmpl"
	LicensesFile = "licenses.txt"
	SwaggerFile = "swagger/example.swagger.json"

	Assets     = assets.AssetFile()
	AssetNames = assets.AssetNames()
	Resources  = resources.AssetFile()

)
