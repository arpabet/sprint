/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintcore

import (
	"go.arpabet.com/glue"
	memstore "go.arpabet.com/store/providers/mem"
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"reflect"
)

type implCacheStoreFactory struct {
	beanName        string
}

func CacheStoreFactory(beanName string) glue.FactoryBean {
	return &implCacheStoreFactory{beanName: beanName}
}

func (t *implCacheStoreFactory) Object() (object interface{}, err error) {

	defer sprintutils.PanicToError(&err)

	return memstore.New(t.beanName), nil
}

func (t *implCacheStoreFactory) ObjectType() reflect.Type {
	return memstore.ObjectType()
}

func (t *implCacheStoreFactory) ObjectName() string {
	return t.beanName
}

func (t *implCacheStoreFactory) Singleton() bool {
	return true
}
