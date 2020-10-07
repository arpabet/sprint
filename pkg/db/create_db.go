/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package db

import (
	"errors"
	"fmt"
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/util"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"os"
	"path/filepath"
)

func HasDatabase(dataDir string) bool {
	if _, err := os.Stat(dataDir); err == nil {
		return true
	}
	return false
}

func CreateDatabase(dataDir string, masterKey string) error {

	fmt.Printf("Create database on folder %s\n", dataDir)

	if _, err := os.Stat(dataDir); err == nil {
		fmt.Printf("Data directory is not empty %s\n", dataDir)
		answer  := util.Prompt("Do you want to delete it? [Y,N]:")
		if answer != "Y" && answer != "y"  {
			return errors.New("operation was canceled by user")
		}
		os.RemoveAll(dataDir)
		err = os.MkdirAll(dataDir, 0777)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Create directory %s\n", dataDir)
	if err := os.MkdirAll(dataDir, 0777); err != nil {
		return err
	}

	keyBytes, err := util.ParseMasterKey(masterKey)
	if err != nil {
		return err
	}

	db, err := OpenDatabase(dataDir, keyBytes)
	if err != nil {
		return err
	}

	defer db.Close()

	return PostInitialize(NewStorageFromDB(db))
}

func OpenDatabase(dataDir string, masterKey []byte) (*badger.DB, error) {

	keyDir := filepath.Join(dataDir, "key")
	valueDir := filepath.Join(dataDir, "value")

	opts := badger.DefaultOptions(dataDir)
	opts.Logger = NewDBLogger()
	opts.Compression = options.ZSTD
	opts.Dir = keyDir
	opts.ValueDir = valueDir
	if app.UseMemoryMap() {
		opts.TableLoadingMode = options.MemoryMap
		opts.ValueLogLoadingMode = options.MemoryMap
	} else {
		opts.TableLoadingMode = options.FileIO
		opts.ValueLogLoadingMode = options.FileIO
	}
	opts.EncryptionKey = masterKey
	opts.EncryptionKeyRotationDuration = app.KeyRotationDuration
	return badger.Open(opts)

}

func PostInitialize(storage app.Storage) error {
	nodeId, err := util.GenerateNodeId()
	if err != nil {
		return err
	}
	println("NodeId:")
	println(nodeId)
	return storage.Put([]byte(app.ConfigPrefix + app.NodeId), []byte(nodeId))
}

