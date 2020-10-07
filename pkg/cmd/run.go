/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/db"
	"github.com/arpabet/sprint/pkg/run"
	"github.com/arpabet/sprint/pkg/util"
	"github.com/pkg/errors"
)

type runCommand struct {
}

func (t *runCommand) Desc() string {
	return "run server"
}

func (t *runCommand) Run(args []string) error {
	app.ParseFlags(args)

	if !db.HasDatabase(app.GetDataFolder()) {
		return errors.Errorf("Database not found in %s", app.GetDataFolder())
	}

	masterKey := util.PromptMasterKey()
	return run.ServerRun(masterKey)
}
