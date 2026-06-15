/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"go.arpabet.com/glue"
	"github.com/pkg/errors"
	"go.arpabet.com/sprint/sprint"
	"go.arpabet.com/sprint/sprintframework/sprintapp"
	"go.arpabet.com/sprint/sprintframework/sprintclient"
	"go.arpabet.com/sprint/sprintframework/sprintcmd"
	"go.arpabet.com/sprint/sprintframework/sprintcore"
	"go.arpabet.com/sprint/sprintframework/sprintserver"
	"log"
	"os"
)

var (
	Version string
	Build   string
)

func doMain() (err error) {

	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			case string:
				err = errors.New(v)
			default:
				err = errors.Errorf("%v", v)
			}
		}
	}()

	glue.Verbose(log.Default())

	beans := []interface{} {
		sprintapp.ApplicationBeans,
		sprintcmd.ApplicationCommands,

		/**
		Those resources and assets are application specific
		 */
		sprintapp.DefaultResources,
		sprintapp.DefaultAssets,
		sprintapp.DefaultGzipAssets,

		glue.Child(sprint.CoreRole,
			sprintcore.CoreServices,
			sprintcore.BoltStoreFactory("config-store"),
			sprintcore.BadgerStoreFactory("secure-store"),
			sprintcore.AutoupdateService(),
			sprintcore.LumberjackFactory(),

			glue.Child(sprint.ServerRole,
				sprintserver.GrpcServerScanner("control-grpc-server"),
				sprintserver.ControlServer(),
				sprintserver.HttpServerFactory("control-gateway-server"),
				//sprintserver.TlsConfigFactory("tls-config"),
				sprintserver.TemplatePage("/", "resources:templates/index.tmpl"),
				),

			glue.Child(sprint.ServerRole,
				sprintserver.HttpServerScanner("redirect-https"),
				sprintserver.RedirectHttpsPage("redirect-https"),
				),
			),
		glue.Child(sprint.ControlClientRole,
			sprintclient.ControlClientBeans,
			//sprintclient.AnyTlsConfigFactory("client-tls-config"),
			),
	}

	return sprintapp.Application("sprint",
		sprintapp.WithVersion(Version),
		sprintapp.WithBuild(Build),
		sprintapp.WithBeans(beans)).
		Run(os.Args[1:])

}

func main() {

	if err := doMain(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
