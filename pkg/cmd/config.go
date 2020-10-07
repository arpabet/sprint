/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/service"
	"github.com/arpabet/sprint/pkg/client"
	"github.com/arpabet/sprint/pkg/util"
	"io"
	"os"
	"fmt"
	"strings"
)

type configCommand struct {
}

func (t *configCommand) Desc() string {
	return "get full config"
}

func (t *configCommand) Run(args []string) error {
	app.ParseFlags(args)
	if err := client.GetConfiguration(os.Stdout); err != nil {
		return getAllInStorage(os.Stdout)
	}
	return nil
}

func getAllInStorage(writer io.StringWriter) error {
	masterKey := util.PromptMasterKey()

	if ctx, err := service.CreateContext(masterKey); err != nil {
		return err
	} else {
		return ctx.MustBean(app.ConfigServiceClass).(app.ConfigService).GetAll(func(key, value string) {
			if strings.Contains(key, "password") {
				value = "******"
			}
			writer.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		})
	}

}


