/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/node"
	"github.com/arpabet/sprint/pkg/service"
	"github.com/arpabet/sprint/pkg/util"
	"github.com/pkg/errors"
)

type getCommand struct {
}

func (t *getCommand) Desc() string {
	return "get config parameter"
}

func (t *getCommand) Run(args []string) error {

	if len(args) < 1 {
		return errors.Errorf("expected one or more arguments: %v", args)
	}

	key := args[0]

	app.ParseFlags(args[1:])

	value, err := node.GetConfig(key)
	if err != nil {
		value, err = getFromStorage(key)
	}
	if err != nil {
		return err
	}

	println(value)
	return nil
}

func getFromStorage(key string) (string, error) {
	masterKey := util.PromptMasterKey()

	if ctx, err := service.CreateContext(masterKey); err != nil {
		return "", err
	} else {
		return ctx.MustBean(app.ConfigServiceClass).(app.ConfigService).Get(key)
	}

}