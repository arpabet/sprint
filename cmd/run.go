/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/template-server/pkg/run"
	"github.com/arpabet/template-server/pkg/util"
)

type runCommand struct {
}

func (t *runCommand) Desc() string {
	return "run server"
}

func (t *runCommand) Run(args []string) error {
	masterKey := util.PromptPassword("Enter master key:")
	return run.ServerRun(false, masterKey)
}
