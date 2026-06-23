/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintcmd

import (
	"fmt"
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/sprint"
	"strings"
)

type implRestartNode struct {
	Application      sprint.Application      `inject:""`
	Context glue.Container `inject:""`
}

func RestartNode() *implRestartNode {
	return &implRestartNode{}
}

func (t *implRestartNode) BeanName() string {
	return "restart"
}

func (t *implRestartNode) Help() string {
	helpText := `
Usage: ./%s restart

	Restarts the application node.

`
	return strings.TrimSpace(fmt.Sprintf(helpText, t.Application.Executable()))
}

func (t *implRestartNode) Synopsis() string {
	return "restart server"
}

func (t *implRestartNode) Run(args []string) error {

	return sprint.DoWithControlClient(t.Context, func(client sprint.ControlClient) error {
		status, err := client.Shutdown(true)
		if err == nil {
			println(status)
		}
		return err
	})

}
