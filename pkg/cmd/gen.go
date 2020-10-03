package cmd

import (
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/util"
)

type genCommand struct {
}

func (t *genCommand) Desc() string {
	return "generate master key"
}

func (t *genCommand) Run(args []string) error {

	app.ParseFlags(args)
	if key, err := util.GenerateMasterKey(); err == nil {
		println(key)
		return nil
	} else {
		return err
	}

}

