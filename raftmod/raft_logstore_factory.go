/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package raftmod

import (
	"go.arpabet.com/glue"
	"go.arpabet.com/store"
	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
	"github.com/pkg/errors"
	"go.arpabet.com/raft-badger"
	"reflect"
)

var LogStoreClass = reflect.TypeOf((*raft.LogStore)(nil)).Elem()

type implRaftLogStoreFactory struct {

	RaftStore     store.ManagedDataStore    `inject:"bean=raft-store"`
	RaftLogPrefix string `value:"raft-store.log-prefix,default=log"`

}

func RaftLogStoreFactory() glue.FactoryBean {
	return &implRaftLogStoreFactory{}
}

func (t *implRaftLogStoreFactory) Object() (object interface{}, err error) {

	defer panicToError(&err)

	db, ok := t.RaftStore.Instance().(*badger.DB)
	if !ok {
		return nil, errors.New("managed data delegate 'raft-store' must have badger backend")
	}

	return raftbadger.NewLogStore(db, []byte(t.RaftLogPrefix)), nil

}

func (t *implRaftLogStoreFactory) ObjectType() reflect.Type {
	return LogStoreClass
}

func (t *implRaftLogStoreFactory) ObjectName() string {
	return "raft-store-log"
}

func (t *implRaftLogStoreFactory) Singleton() bool {
	return true
}
