/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintcmd

import (
	"fmt"
	"go.arpabet.com/glue"
	"github.com/pkg/errors"
	"go.arpabet.com/sprint/sprint"
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type implStopNode struct {
	Application      sprint.Application      `inject`
	ApplicationFlags sprint.ApplicationFlags `inject`
	Context          glue.Container         `inject`

	RunDir           string       `value:"application.run.dir,default="`
}

func StopNode() *implStopNode {
	return &implStopNode{}
}

func (t *implStopNode) Run(args []string) error {

	err := sprint.DoWithControlClient(t.Context, func(client sprint.ControlClient) error {
		status, err := client.Shutdown(false)
		if err == nil {
			println(status)
		}
		return err
	})

	if err != nil {
		return t.KillServer()
	}

	return nil
}

func (t *implStopNode) KillServer() error {

	runDir := t.RunDir
	if runDir == "" {
		runDir = filepath.Join(t.Application.ApplicationDir(), "run")
	}
	pidFile := filepath.Join(runDir, fmt.Sprintf("%s.pid", t.getNodeName()))

	blob, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return err
	}

	pid := string(blob)

	if _, err := strconv.Atoi(pid); err != nil {
		return errors.Errorf("Invalid pid %s, %v", pid, err)
	}

	cmd := exec.Command("kill", "-2", pid)
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := os.Remove(pidFile); err != nil {
		return errors.Errorf("Can not remove file %s, %v", pidFile, err)
	}

	return cmd.Wait()

}

func (t *implStopNode) getNodeName() string {
	return sprintutils.AppendNodeSequence(t.Application.Name(), t.ApplicationFlags.Node())
}