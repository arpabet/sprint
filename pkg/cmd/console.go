package cmd

import (
	"github.com/arpabet/sprint/pkg/node"
	"os"
)

type consoleCommand struct {
}

func (t *consoleCommand) Desc() string {
	return "database console"
}

func (t *consoleCommand) Run(args []string) error {
	return node.DatabaseConsole(os.Stdout, os.Stderr)
}

