/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/db"
	"github.com/arpabet/sprint/pkg/util"
	"github.com/pkg/errors"
)

type startCommand struct {
}

func (t *startCommand) Desc() string {
	return "start server"
}

func (t *startCommand) Run(args []string) error {
	app.ParseFlags(args)

	if !db.HasDatabase(app.GetDataFolder()) {
		return errors.Errorf("Database not found in %s", app.GetDataFolder())
	}

	masterKey := util.PromptMasterKey()

	if _, err := util.ParseMasterKey(masterKey); err != nil {
		return err
	}

	return util.StartServer(masterKey)
}
