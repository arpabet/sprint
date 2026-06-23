/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintclient

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"path/filepath"
	"reflect"

	"go.arpabet.com/glue"
	"go.arpabet.com/properties"
	"go.arpabet.com/sprint/sprint"
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"golang.org/x/xerrors"
)

var (
	CertFile = "client.crt"
	KeyFile  = "client.key"
)

type tlsConfigFactory struct {
	Application sprint.Application `inject:""`
	Properties  glue.Properties    `inject:""`

	CompanyName string `value:"application.company,default=sprint"`

	beanName string
}

func TlsConfigFactory(beanName string) glue.FactoryBean {
	return &tlsConfigFactory{beanName: beanName}
}

func (t *tlsConfigFactory) Object() (object interface{}, err error) {

	defer sprintutils.PanicToError(&err)

	appDir := properties.Locate(t.CompanyName).GetDir(t.Application.Name())

	certFile := filepath.Join(appDir, CertFile)
	keyFile := filepath.Join(appDir, KeyFile)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, xerrors.Errorf("LoadX509KeyPair for implControlClient SSL from %s and %s failed, %v", certFile, keyFile, err)
	}

	insecure := t.Properties.GetBool(fmt.Sprintf("%s.insecure", t.beanName), false)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: insecure,
		Rand:               rand.Reader,
	}

	tlsConfig.NextProtos = appendH2ToNextProtos(tlsConfig.NextProtos)
	return tlsConfig, err
}

func (t *tlsConfigFactory) ObjectType() reflect.Type {
	return sprint.TlsConfigClass
}

func (t *tlsConfigFactory) ObjectName() string {
	return t.beanName
}

func (t *tlsConfigFactory) Singleton() bool {
	return true
}
