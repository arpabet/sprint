/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package cmd

import (
	"github.com/arpabet/templateserv/pkg/app"
)


type licensesCommand struct {

}
func (t *licensesCommand) Desc() string {
	return "show all licenses"
}

func (t *licensesCommand) Run(args []string) error {
	print(app.GetLicenses())
	return nil
}



