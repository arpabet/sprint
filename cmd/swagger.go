/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package cmd

import (
	"github.com/arpabet/templateserv/pkg/app"
)

type swaggerCommand struct {
}

func (t *swaggerCommand) Desc() string {
	return "swagger description"
}

func (t *swaggerCommand) Run(args []string) error {
	print(app.GetSwagger())
	return nil
}
