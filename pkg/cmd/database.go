package cmd

import (
	"github.com/arpabet/sprint/pkg/node"
	"os"
)

type databaseCommand struct {
}

func (t *databaseCommand) Desc() string {
	return "database console"
}

func (t *databaseCommand) Run(args []string) error {
	return node.DatabaseConsole(os.Stdout, os.Stderr)
}

