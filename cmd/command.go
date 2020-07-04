/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
*/

package cmd

import (
	"flag"
	"fmt"
	"github.com/arpabet/template-server/pkg/constants"
)

type commandFace interface {

	Run(args []string) error

	Desc() string

}

var allCommands = map[string]commandFace {

	"version": &versionCommand{},

	"create": &createCommand{},

	"start": &startCommand{},

	"run": &runCommand{},

	"status": &statusCommand{},

	"stop": &stopCommand{},

	"ssl": &sslCommand{},

	"licenses": &licensesCommand{},

	"help": &helpCommand{},

}

func preprocessArgs(args []string) []string {

	if len(args) == 1 && (args[0] == "-v" || args[0] == "-version" || args[0] == "--version") {
		return []string{"version"}
	}

	return args
}

func printUsage() {

	fmt.Printf("Usage: %s [command]\n", constants.ExecutableName)

	for name, command := range allCommands {
		fmt.Printf("    %s - %s\n", name, command.Desc())
	}

	fmt.Println("Flags:")
	flag.PrintDefaults()

}

func Run(args []string) int {

	args = preprocessArgs(args)

	if len(args) >= 1 {

		cmd := args[0]

		if inst, ok := allCommands[cmd]; ok {

			if err := inst.Run(args[1:]); err != nil {
				fmt.Printf("Error: %v\n", err)
				return 1
			}
			return 0

		} else {
			fmt.Printf("Invalid command: %s\n", cmd)
			printUsage()
			return 1
		}

	} else {
		printUsage()
		return 0
	}

	return 0
}



