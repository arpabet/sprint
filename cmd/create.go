/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package cmd

import (
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/db"
	"github.com/arpabet/templateserv/pkg/util"
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
		return db.CreateDatabase(app.GetDataFolder(), key)
	} else {
		return err
	}
}
