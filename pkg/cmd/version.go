/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package cmd

import (
	"fmt"
	"github.com/arpabet/templateserv/pkg/app"
)


type versionCommand struct {
}

func (t *versionCommand) Desc() string {
	return "show version"
}

func (t *versionCommand) Run(args []string) error {

	fmt.Printf("%s [Version %s, Build %s]\n", app.ApplicationName, app.Version, app.Build)
	fmt.Printf("%s\n", app.Copyright)
	return nil
}
