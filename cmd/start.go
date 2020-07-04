/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import "github.com/arpabet/template-server/pkg/util"

type startCommand struct {
}

func (t *startCommand) Desc() string {
	return "start server"
}

func (t *startCommand) Run(args []string) error {
	masterKey := util.PromptPassword("Enter master key:")

	if _, err := util.ParseMasterKey(masterKey); err != nil {
		return err
	}

	return util.StartServer(masterKey)
}
