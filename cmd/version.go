/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import (
	"fmt"
	"github.com/arpabet/template-server/pkg/constants"
)


type versionCommand struct {
}

func (t *versionCommand) Desc() string {
	return "show version"
}

func (t *versionCommand) Run(args []string) error {

	appInfo := constants.GetAppInfo()
	fmt.Printf("%s [Version %s, Build %s]\n", constants.ApplicationName, appInfo.Version, appInfo.Build)
	fmt.Printf("%s\n", constants.Copyright)
	return nil
}
