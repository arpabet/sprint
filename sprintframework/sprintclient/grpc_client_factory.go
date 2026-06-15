/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintclient

import (
	"crypto/tls"
	"fmt"
	"go.arpabet.com/glue"
	"github.com/pkg/errors"
	"go.arpabet.com/sprint/sprint"
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"reflect"
	"strings"
)

type implGrpcClientFactory struct {

	Application       sprint.Application       `inject`
	ApplicationFlags  sprint.ApplicationFlags  `inject`
	Properties        glue.Properties          `inject`
	TlsConfig         *tls.Config              `inject:"optional"`

	beanName string
}

func GrpcClientFactory(beanName string) glue.FactoryBean {
	return &implGrpcClientFactory{
		beanName: beanName,
	}
}

func (t *implGrpcClientFactory) Object() (object interface{}, err error) {

	defer sprintutils.PanicToError(&err)

	// try to get normal property
	connectAddr := t.Properties.GetString(fmt.Sprintf("%s.connect-address", t.beanName), "")
	if connectAddr == "" {
		// try to lookup from server address
		serverBean := strings.ReplaceAll(t.beanName, "client", "server")
		grpcListenAddr := t.Properties.GetString(fmt.Sprintf("%s.bind-address", serverBean), "")
		if grpcListenAddr == "" {
			return nil, errors.Errorf("property '%s.connect-address' is not found and property '%s.bind-address' is not found too'", t.beanName, serverBean)
		}
		connectAddr = t.getConnectFromBindAddress(grpcListenAddr)
	}

	tcpAddr, err := sprintutils.ParseAndAdjustTCPAddr(connectAddr, t.ApplicationFlags.Node())
	if err != nil {
		return
	}
	connectAddr = fmt.Sprintf("%s:%d", tcpAddr.IP.String(), tcpAddr.Port)

	return t.doDial(connectAddr)
}

func (t *implGrpcClientFactory) ObjectType() reflect.Type {
	return sprint.GrpcClientConnClass
}

func (t *implGrpcClientFactory) ObjectName() string {
	return t.beanName
}

func (t *implGrpcClientFactory) Singleton() bool {
	return true
}

func (t *implGrpcClientFactory) getConnectFromBindAddress(listenAddr string) string {
	if strings.HasPrefix(listenAddr, "0.0.0.0:") {
		return "127.0.0.1" + listenAddr[7:]
	}
	if strings.HasPrefix(listenAddr, ":") {
		return "127.0.0.1" + listenAddr
	}
	return listenAddr
}

func (t *implGrpcClientFactory) getTransportCreds() credentials.TransportCredentials {
	if t.TlsConfig != nil {
		return credentials.NewTLS(t.TlsConfig)
	} else {
		return insecure.NewCredentials()
	}
}

func (t *implGrpcClientFactory) doDial(connectAddr string) (*grpc.ClientConn, error) {

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(t.getTransportCreds()))

	maxMessageSize := t.Properties.GetInt(fmt.Sprintf("%s.max-message-size", t.beanName), 0)
	if maxMessageSize != 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMessageSize)))
	}

	authToken := t.Properties.GetString("application.auth", "")
	if authToken != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(&tokenAuth{token: authToken}))
	}

	return grpc.Dial(connectAddr, opts...)
}
