/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/db"
	"github.com/arpabet/sprint/pkg/util"
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
