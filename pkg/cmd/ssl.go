/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package cmd

import (
	"fmt"
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/db"
	"github.com/arpabet/sprint/pkg/util"
	"os"
)

type sslCommand struct {
}


func (t *sslCommand) Desc() string {
	return "import SSL certificates"
}

func (t *sslCommand) Run(args []string) error {

	println("Set SSL Certificates")

	sslDir := app.DefaultSSLFolder
	if len(args) >= 1 {
		sslDir = args[0]
		args = args[1:]
	}

	app.ParseFlags(args)

	masterKey := util.PromptMasterKey()

	storage, err := db.NewStorage(app.GetDataFolder(), masterKey)
	if err != nil {
		return err
	}
	defer storage.Close()

	if _, err := os.Stat(sslDir); os.IsNotExist(err) {
		fmt.Printf("SSL folder %s is not exist\n", sslDir)
		return util.PromptCertificates(storage)
	} else {
		println("Import ssl certificates from folder: " + sslDir)
		return util.ImportCertificates(storage, sslDir)
	}


}
