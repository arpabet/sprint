/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package db

import (
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/util"
	"github.com/consensusdb/value"
	"github.com/dgraph-io/badger/v2"
	"github.com/pkg/errors"
)


type storage struct {
	db     *badger.DB
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

func NewStorageFromDB(db *badger.DB) app.Storage {
	return &storage {db}
}

func (t* storage) Close() error {
	return t.db.Close()
}


func (t* storage) Get(key []byte, required bool) ([]byte, error) {
	return t.getImpl(key, required)
}

func (t* storage) GetValue(key []byte, required bool) (value.Value, error) {
	content, err := t.Get(key, required)
	if err != nil || content == nil {
		return nil, err
	}
	return value.Unpack(content, false)
}

func (t* storage) Put(key []byte, content []byte) error {

	txn := t.db.NewTransaction(true)
	defer txn.Discard()

	entry := &badger.Entry{ Key: key, Value: content, UserMeta: byte(0x0) }
	err := txn.SetEntry(entry)

	if err != nil {
		return errors.Errorf("Put entry error: %v", err)
	}

	return txn.Commit()

}

func (t* storage) PutValue(key []byte, val value.Value) error {
	content, err := value.Pack(val)
	if err != nil {
		return err
	}
	return t.Put(key, content)
}

func (t* storage) Remove(key []byte) error {

	txn := t.db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Delete(key)

	if err != nil {
		return errors.Errorf("Delete entry error: %v", err)
	}
	return txn.Commit()
}

func (t* storage) getImpl(key []byte, required bool) ([]byte, error) {

	txn := t.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(key)
	if err != nil {

		if err == badger.ErrKeyNotFound && !required {
			return nil, nil
		} else {
			return nil, errors.Errorf("DB required key '%s' not found, %v", key, err)
		}

	}

	data, err := item.ValueCopy(nil)
	if err != nil {
		return nil, errors.Errorf("fetch value failed: %v", err)
	}

	return data, nil
}

func (t* storage) Enumerate(prefix, key []byte, batchSize int, cb func(key, value []byte) bool) error {

	options := badger.IteratorOptions{
		PrefetchValues: true,
		PrefetchSize:   batchSize,
		Reverse:        true,
		AllVersions:    false,
		Prefix: 		prefix,
	}

	txn := t.db.NewTransaction(false)
	defer txn.Discard()

	iter := txn.NewIterator(options)
	defer iter.Close()

	iter.Seek(key)

	for ;iter.Valid(); iter.Next() {

		item := iter.Item()
		key := item.Key()
		value, err := item.ValueCopy(nil)
		if err != nil {
			return errors.Errorf("db: failed to copy value for key %v", key)
		}
		if !cb(key, value) {
			break
		}

	}

	return nil
}