/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/template-server/pkg/client"
)

type statusCommand struct {
}

func (t *statusCommand) Desc() string {
	return "server status"
}

func (t *statusCommand) Run(args []string) error {
	status, err := client.RequestStatus()
	if err != nil {
		return err
	}

	println(status.String())
	return nil
}
