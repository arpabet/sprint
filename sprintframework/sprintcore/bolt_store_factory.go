/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintcore

import (
	"fmt"
	"go.arpabet.com/glue"
	boltstore "go.arpabet.com/store/providers/bolt"
	"go.arpabet.com/sprint/sprint"
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"os"
	"path/filepath"
	"reflect"
)

type implBoltStoreFactory struct {
	beanName        string

	Application      sprint.Application      `inject`
	ApplicationFlags sprint.ApplicationFlags `inject`
	Properties       glue.Properties         `inject`

	DataDir           string       `value:"application.data.dir,default="`
	DataDirPerm       os.FileMode  `value:"application.perm.data.dir,default=-rwxrwx---"`
	DataFilePerm      os.FileMode  `value:"application.perm.data.file,default=-rw-rw-r--"`
}

func BoltStoreFactory(beanName string) glue.FactoryBean {
	return &implBoltStoreFactory{beanName: beanName}
}

func (t *implBoltStoreFactory) Object() (object interface{}, err error) {

	defer sprintutils.PanicToError(&err)

	dataDir := t.DataDir
	if dataDir == "" {
		dataDir = filepath.Join(t.Application.ApplicationDir(), "db")

		if err := sprintutils.CreateDirIfNeeded(dataDir, t.DataDirPerm); err != nil {
			return nil, err
		}

		dataDir = filepath.Join(dataDir, t.getNodeName())
	}
	if err := sprintutils.CreateDirIfNeeded(dataDir, t.DataDirPerm); err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s.db", t.beanName)
	dataFile := filepath.Join(dataDir, fileName)

	return boltstore.New(t.beanName, dataFile, t.DataFilePerm)
}

func (t *implBoltStoreFactory) ObjectType() reflect.Type {
	return boltstore.ObjectType()
}

func (t *implBoltStoreFactory) ObjectName() string {
	return t.beanName
}

func (t *implBoltStoreFactory) Singleton() bool {
	return true
}

func (t *implBoltStoreFactory) getNodeName() string {
	return sprintutils.AppendNodeSequence(t.Application.Name(), t.ApplicationFlags.Node())
}