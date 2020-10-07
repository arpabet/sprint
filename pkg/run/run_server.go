/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package run

import (
	"github.com/arpabet/sprint/pkg/service"
)

func ServerRun(masterKey string) error {

	ctx, err := service.CreateContext(masterKey)
	if err != nil {
		return err
	}
	defer ctx.Close()

	srv := NewServerImpl(ctx)
	if err := ctx.Inject(srv); err != nil {
		return err
	}

	return srv.Run(masterKey)
}
