/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintclient

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/sprint"
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"reflect"
)

type implAnyTlsConfigFactory struct {
	Properties    glue.Properties  `inject:""`
	beanName string
}

func AnyTlsConfigFactory(beanName string) glue.FactoryBean {
	return &implAnyTlsConfigFactory{beanName: beanName}
}

func (t *implAnyTlsConfigFactory) Object() (object interface{}, err error) {

	defer sprintutils.PanicToError(&err)

	insecure := t.Properties.GetBool(fmt.Sprintf("%s.insecure", t.beanName), false)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecure,
		Rand:               rand.Reader,
	}

	tlsConfig.NextProtos = appendH2ToNextProtos(tlsConfig.NextProtos)
	return tlsConfig, nil
}

func (t *implAnyTlsConfigFactory) ObjectType() reflect.Type {
	return sprint.TlsConfigClass
}

func (t *implAnyTlsConfigFactory) ObjectName() string {
	return t.beanName
}

func (t *implAnyTlsConfigFactory) Singleton() bool {
	return true
}
