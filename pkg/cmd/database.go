package cmd

import (
	"github.com/arpabet/sprint/pkg/client"
	"os"
)

type databaseCommand struct {
}

func (t *databaseCommand) Desc() string {
	return "database console"
}

func (t *databaseCommand) Run(args []string) error {
	return client.DatabaseConsole(os.Stdout, os.Stderr)
}

