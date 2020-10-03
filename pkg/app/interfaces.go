/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	"github.com/arpabet/templateserv/pkg/pb"
	"github.com/arpabet/timeuuid"
	"github.com/arpabet/value"
	"reflect"
)

var StorageClass = reflect.TypeOf((*Storage)(nil)).Elem()
type Storage interface {

	Get(key []byte, required bool) ([]byte, error)

	GetValue(key []byte, required bool) (value.Value, error)

	Put(key, content []byte) error

	PutValue(key []byte, value value.Value) error

	Remove(key []byte) error

	Enumerate(prefix, key []byte, batchSize int, cb func(key, value []byte) bool) error

	Close() error

}

var ConfigServiceClass = reflect.TypeOf((*ConfigService)(nil)).Elem()
type ConfigService interface {

	Get(key string) (string, error)

	GetWithDefault(key, defaultValue string) (string, error)

	GetAll(func(key, value string)) error

	GetBool(key string) (bool, error)

	GetInt(key string, defaultValue int) (int, error)

	Set(key, value string) error

}

var NodeServiceClass = reflect.TypeOf((*NodeService)(nil)).Elem()
type NodeService interface {

	NodeId() uint64

	NodeIdHex() string

	Issue() timeuuid.UUID

	Parse(timeuuid.UUID) (timestampMillis int64, nodeId int64, clock int)

}

type DatabaseConsoleStream interface {

	Send(*pb.DatabaseResponse) error

	Recv() (*pb.DatabaseRequest, error)

}


type Record struct {
	Key   []byte
	Value []byte
}

var DatabaseServiceClass = reflect.TypeOf((*DatabaseService)(nil)).Elem()
type DatabaseService interface {

	Execute(query string, cb func(Record) bool) error

	Console(stream DatabaseConsoleStream) error

}

