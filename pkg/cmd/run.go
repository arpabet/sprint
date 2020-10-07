/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/run"
	"github.com/arpabet/sprint/pkg/util"
)

type runCommand struct {
}

func (t *runCommand) Desc() string {
	return "run server"
}

func (t *runCommand) Run(args []string) error {
	app.ParseFlags(args)
	masterKey := util.PromptMasterKey()
	return run.ServerRun(masterKey)
}
