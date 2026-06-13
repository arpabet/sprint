/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package sprintcore

import (
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/sprint"
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"go.uber.org/zap"
	"reflect"
)

type implHCLogFactory struct {
	Log              *zap.Logger             `inject`
}

func HCLogFactory() glue.FactoryBean {
	return &implHCLogFactory{}
}

func (t *implHCLogFactory) Object() (object interface{}, err error) {

	defer sprintutils.PanicToError(&err)

	return newHCLogAdapter(t.Log), nil
}

func (t *implHCLogFactory) ObjectType() reflect.Type {
	return sprint.HCLogClass
}

func (t *implHCLogFactory) ObjectName() string {
	return "hclog_logger"
}

func (t *implHCLogFactory) Singleton() bool {
	return true
}

