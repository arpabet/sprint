/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package util

import (
	"io/ioutil"
	"net/http"
	"os"
)

func CopyFile(src string, dst string, perm os.FileMode) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, data, perm)
	if err != nil {
		return err
	}
	return nil
}

type combinedFileSystem struct {
	list  []http.FileSystem
}

func (t combinedFileSystem) Open(name string) (file http.File, err error) {
	for _, fs := range t.list {
		if fs != nil {
			file, err = fs.Open(name)
			if err == nil {
				return file, nil
			}
		}
	}
	return
}

func CombineFileSystems(fs... http.FileSystem) http.FileSystem {
	return &combinedFileSystem{list: fs}
}


func IsFileLocked(filePath string) bool {
	if file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_EXCL, 0); err != nil {
		return true
	} else {
		file.Close()
		return false
	}
}