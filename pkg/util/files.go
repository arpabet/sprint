/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package util

import (
	"io/ioutil"
	"os"
)

func copyFile(src string, dst string, perm os.FileMode) error {
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

