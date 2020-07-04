/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */

package cmd

import (
	"github.com/arpabet/template-server/pkg/constants"
	"github.com/arpabet/template-server/pkg/db"
	"github.com/arpabet/template-server/pkg/util"
)

type createCommand struct {
}

func (t *createCommand) Desc() string {
	return "create database"
}

func (t *createCommand) Run(args []string) error {
	println("Generated master key:")
	if key, err := util.GenerateMasterKey(); err == nil {
		println(key)
		return db.CreateDatabase(constants.GetDatabaseFolder(), key)
	} else {
		return err
	}
}
