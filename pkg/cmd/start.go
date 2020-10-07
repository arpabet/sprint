/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/util"
)

type startCommand struct {
}

func (t *startCommand) Desc() string {
	return "start server"
}

func (t *startCommand) Run(args []string) error {
	app.ParseFlags(args)

	masterKey := util.PromptMasterKey()

	if _, err := util.ParseMasterKey(masterKey); err != nil {
		return err
	}

	return util.StartServer(masterKey)
}
