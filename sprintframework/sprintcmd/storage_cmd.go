/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintcmd

import (
	"fmt"
	"go.arpabet.com/glue"
	"github.com/pkg/errors"
	"go.arpabet.com/sprint/sprint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

type implStorageCommand struct {
	Application      sprint.Application      `inject`
	Context          glue.Container           `inject`
}

type coreStorageContext struct {
	StorageService sprint.StorageService `inject`
}

func StorageCommand() sprint.Command {
	return &implStorageCommand{}
}

func (t *implStorageCommand) BeanName() string {
	return "storage"
}

func (t *implStorageCommand) Help() string {
	helpText := `
Usage: ./%s storage [command]

	Provides management functionality for the internal storage(s).

Commands:

  console                  Runs interactive storage console.

  list                     Lists available internal storages.

  dump                     Dumps the internal storage to the file.

  restore                  Restore internal storage from the file.

  compact                  Runs compaction background process on internal storage.

  drop                     Non-reversible operation of deleting all records in the internal storage.

  clean                    Clean the storage.

`
	return strings.TrimSpace(fmt.Sprintf(helpText, t.Application.Executable()))
}

func (t *implStorageCommand) Synopsis() string {
	return "storage management commands: [console, list, dump, restore, compact, drop, clean]"
}

func (t *implStorageCommand) Run(args []string) error {

	if len(args) < 1 {
		return errors.Errorf("storage needs command: %s", t.Synopsis())
	}

	cmd := args[0]
	args = args[1:]

	err := sprint.DoWithControlClient(t.Context, func(client sprint.ControlClient) error {
		if cmd == "console" {
			return client.StorageConsole(os.Stdout, os.Stderr)
		} else {
			output, err := client.StorageCommand(cmd, args)
			if err != nil {
				return err
			}
			println(output)
			return nil
		}
	})
	if err == nil {
		return nil
	}
	if status.Code(err) != codes.Unavailable {
		return err
	}

	c := new(coreStorageContext)
	return doInCore(t.Context, c, func(core glue.Container) error {
		if cmd == "console" {
			return c.StorageService.LocalConsole(os.Stdout, os.Stderr)
		} else {
			content, err :=  c.StorageService.ExecuteCommand(cmd, args)
			if err != nil {
				return err
			}
			println(content)
			return nil
		}
	})

}



