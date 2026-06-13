/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package sprintcmd

import (
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/sprint"
)

type implStatusNode struct {
	Application sprint.Application `inject`
	Context glue.Container `inject`
}

func StatusNode() *implStatusNode {
	return &implStatusNode{}
}

func (t *implStatusNode) Run(args []string) error {

	return sprint.DoWithControlClient(t.Context, func(client sprint.ControlClient) error {
		status, err := client.Status()
		if err == nil {
			println(status)
		}
		return err
	})

}
