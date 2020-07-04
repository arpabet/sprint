/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package db

import (
	"github.com/arpabet/template-server/pkg/app"
	"github.com/arpabet/template-server/pkg/util"
	"github.com/dgraph-io/badger/v2"
	"github.com/pkg/errors"
)

type storage struct {
	db        			*badger.DB
}

func NewStorage(dataDir string, masterKey string) (app.Storage, error) {

	keyBytes, err := util.ParseMasterKey(masterKey)
	if err != nil {
		return nil, err
	}

	db, err := OpenDatabase(dataDir, keyBytes)
	if err != nil {
		return nil, err
	}

	return &storage {db}, nil
}

func (t* storage) Close() error {
	return t.db.Close()
}

func (t* storage) Find(key string) ([]byte, error) {
	return t.getImpl(key, false)
}

func (t* storage) Get(key string) ([]byte, error) {
	return t.getImpl(key, true)
}

func (t* storage) Put(key string, content []byte) error {

	txn := t.db.NewTransaction(true)
	defer txn.Discard()

	entry := &badger.Entry{ Key: []byte(key), Value: content, UserMeta: byte(0x0) }
	err := txn.SetEntry(entry)

	if err != nil {
		return errors.Errorf("Put entry error: %v", err)
	}

	return txn.Commit()

}

func (t* storage) Remove(key string) error {

	txn := t.db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Delete([]byte(key))

	if err != nil {
		return errors.Errorf("Delete entry error: %v", err)
	}
	return txn.Commit()
}

func (t* storage) getImpl(key string, required bool) ([]byte, error) {

	txn := t.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(key))
	if err != nil {

		if err == badger.ErrKeyNotFound && !required {
			return nil, nil
		} else {
			return nil, err
		}

	}

	data, err := item.ValueCopy(nil)
	if err != nil {
		return nil, errors.Errorf("fetch value failed: %v", err)
	}

	return data, nil
}
