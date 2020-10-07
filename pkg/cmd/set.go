package cmd

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/node"
	"github.com/arpabet/sprint/pkg/service"
	"github.com/arpabet/sprint/pkg/util"
	"github.com/pkg/errors"
)

type setCommand struct {
}

func (t *setCommand) Desc() string {
	return "set config parameter"
}

func (t *setCommand) Run(args []string) error {

	if len(args) < 2 {
		return errors.Errorf("expected two or more arguments: %v", args)
	}

	key := args[0]
	value := args[1]

	app.ParseFlags(args[2:])

	status, err := node.SetConfig(key, value)
	if err != nil {
		status, err = setToStorage(key, value)
	}
	if err != nil {
		return err
	}

	println(status)
	return nil
}


func setToStorage(key, value string) (string, error) {
	masterKey := util.PromptMasterKey()

	if ctx, err := service.CreateContext(masterKey); err != nil {
		return "", err
	} else if err := ctx.MustBean(app.ConfigServiceClass).(app.ConfigService).Set(key, value); err != nil {
		return "", err
	} else {
		return "OK", nil
	}

}