/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintserver

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/cert"
	"go.arpabet.com/sprint"
	"go.arpabet.com/sprint/framework/sprintutils"
	"reflect"
)

type implTlsConfigFactory struct {

	Properties     glue.Properties      `inject:""`
	NodeService    sprint.NodeService   `inject:""`

	CertificateManager cert.CertificateManager `inject:"optional"`

	beanName          string
}

func TlsConfigFactory(beanName string) glue.FactoryBean {
	return &implTlsConfigFactory{beanName: beanName}
}

func (t *implTlsConfigFactory) Object() (obj interface{}, err error) {

	defer sprintutils.PanicToError(&err)

	insecure := t.Properties.GetBool(fmt.Sprintf("%s.insecure", t.beanName), false)

	tlsConfig := &tls.Config{
		Rand:         rand.Reader,
		InsecureSkipVerify: insecure,
	}

	if t.CertificateManager != nil {
		tlsConfig.GetCertificate = t.CertificateManager.GetCertificate
	}

	tlsConfig.NextProtos = AppendH2ToNextProtos(tlsConfig.NextProtos)
	return tlsConfig, nil
}

func (t *implTlsConfigFactory) ObjectType() reflect.Type {
	return sprint.TlsConfigClass
}

func (t *implTlsConfigFactory) ObjectName() string {
	return t.beanName
}

func (t *implTlsConfigFactory) Singleton() bool {
	return true
}


