/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	"encoding/base64"
	"os"
	"time"
)

var (

	Version   string
	Build     string

	MasterKey = os.Getenv("TEMPLATE_MASTER_KEY")
	ClassOfService = os.Getenv("COS")

	ExecutableName = "templateserv"
	ExecutablePID = ExecutableName + ".pid"
	ExecutableLog = ExecutableName + ".log"

	ApplicationName = "TemplateServer"
	PackageName = "github.com/arpabet/templateserv"

	Copyright = "Copyright (C) 2020-present Arpabet Inc. All rights reserved."

	DataDirPerm = os.FileMode(0700)

	KeySize = 32  // 256-bit AES key

	NodeIdBits = 48
	NodeIdSize = NodeIdBits / 8

	Encoding = base64.RawURLEncoding

	PlainContentType = "text/plain"

	IsProd = ClassOfService == "prod"
	IsDev = ClassOfService == "dev" || ClassOfService == ""
)

// ones a week change a key
const KeyRotationDuration = time.Hour * 24 * 7

var (
	statusEndpoint = "https://%s/api/status"
	stopEndpoint = "https://%s/api/stop"
	setConfigEndpoint = "https://%s/api/config"
	getConfigEndpoint = "https://%s/api/config/%s"
)

var (

	DefaultSSLFolder = "ssl"

	EventPrefix = byte('#')
	ConfigPrefix = "config:"

	NodeId = "node.id"

	ListenTlsAddress = "listen.tls.address"
	DefaultTlsAddress = "127.0.0.1:8443"

	ListenGrpcAddress = "listen.grpc.address"        			 // if empty then do not run gRPC server
	ListenGrpcGatewayAddress = "listen.grpc.gateway.address"     // if empty then do not run gRPC gateway server

)