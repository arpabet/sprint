/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package app

type Storage interface {

	Find(key string) ([]byte, error)

	Get(key string) ([]byte, error)

	Put(key string, content []byte) error

	Remove(key string) error

	Close() error

}
