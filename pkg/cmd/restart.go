/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/node"
)

type restartCommand struct {
}

func (t *restartCommand) Desc() string {
	return "restart server"
}

func (t *restartCommand) Run(args []string) error {
	app.ParseFlags(args)

	if status, err := node.Shutdown(true); err == nil {
		println(status)
		return nil
	} else {
		return err
	}
}
