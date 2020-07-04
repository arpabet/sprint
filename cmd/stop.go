/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import "github.com/arpabet/template-server/pkg/util"

type stopCommand struct {
}

func (t *stopCommand) Desc() string {
	return "stop server"
}

func (t *stopCommand) Run(args []string) error {
	return util.StopServer()
}
