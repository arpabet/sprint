/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

type helpCommand struct {

}
func (t *helpCommand) Desc() string {
	return "help command"
}

func (t *helpCommand) Run(args []string) error {
	printUsage()
	return nil
}
