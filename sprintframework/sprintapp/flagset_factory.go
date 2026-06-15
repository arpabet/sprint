/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintapp

import (
	"flag"
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/sprint"
	"reflect"
)

type implFlagSetFactory struct {
	Registrars []sprint.FlagSetRegistrar `inject`
}

func FlagSetFactory() glue.FactoryBean {
	return &implFlagSetFactory{}
}

func (t *implFlagSetFactory) Object() (interface{}, error) {
	fs := flag.NewFlagSet("sprint", flag.ContinueOnError)
	for _, reg := range t.Registrars {
		reg.RegisterFlags(fs)
	}
	return fs, nil
}

func (t *implFlagSetFactory) ObjectType() reflect.Type {
	return sprint.FlagSetClass
}

func (t *implFlagSetFactory) ObjectName() string {
	return ""
}

func (t *implFlagSetFactory) Singleton() bool {
	return true
}
