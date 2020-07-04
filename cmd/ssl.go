/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import (
	"fmt"
	"github.com/arpabet/template-server/pkg/constants"
	"github.com/arpabet/template-server/pkg/db"
	"github.com/arpabet/template-server/pkg/util"
	"os"
)

type sslCommand struct {
}


func (t *sslCommand) Desc() string {
	return "import SSL certificates"
}

func (t *sslCommand) Run(args []string) error {

	println("Set SSL Certificates")

	masterKey := util.PromptPassword("Enter master key:")

	storage, err := db.NewStorage(constants.GetDatabaseFolder(), masterKey)
	if err != nil {
		return err
	}
	defer storage.Close()

	sslDir := constants.GetSSLFolder()

	if _, err := os.Stat(sslDir); os.IsNotExist(err) {
		fmt.Printf("SSL folder %s is not exist\n", sslDir)
		return util.PromptCertificates(storage)
	} else {
		return util.ImportCertificates(storage, sslDir)
	}


}
