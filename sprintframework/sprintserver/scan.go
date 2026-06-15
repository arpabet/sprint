/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintserver

import (
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/sprint"
	"google.golang.org/grpc"
	"net/http"
)

type grpcServerScanner struct {
	beanName string
	scan     []interface{}
}

func GrpcServerScanner(beanName string, scan... interface{}) glue.Scanner {
	return &grpcServerScanner{
		beanName: beanName,
		scan:     scan,
	}
}

func (t *grpcServerScanner) ScannerBeans() []interface{} {
	beans := []interface{}{
		AuthorizationMiddleware(),
		GrpcServerFactory(t.beanName),
		&struct {
			// make them visible
			Servers     []sprint.Server `inject:"optional"`
			GrpcServers []*grpc.Server  `inject:"optional"`
			HttpServers []*http.Server  `inject:"optional"`
		}{},
	}
	return append(beans, t.scan...)
}

type httpServerScanner struct {
	beanName string
	scan     []interface{}
}

func HttpServerScanner(beanName string, scan... interface{}) glue.Scanner {
	return &httpServerScanner{
		beanName: beanName,
		scan:     scan,
	}
}

func (t *httpServerScanner) ScannerBeans() []interface{} {
	beans := []interface{}{
		HttpServerFactory(t.beanName),
		&struct {
			// make them visible
			Servers     []sprint.Server `inject:"optional"`
			HttpServers []*http.Server `inject`
		}{},
	}
	return append(beans, t.scan...)
}

