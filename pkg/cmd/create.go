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
	masterKey := util.PromptMasterKey()
	return db.CreateDatabase(app.GetDataFolder(), masterKey)
}
