/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package main

import (
	"github.com/arpabet/template-server/cmd"
	"github.com/arpabet/template-server/pkg/constants"
	"github.com/arpabet/template-server/pkg/app"
	"math/rand"
	"os"
	"time"
)

var (
	Version   string
	Built     string
)

func main() {

	constants.ParseFlags()

	rand.Seed(time.Now().UnixNano())

	constants.SetAppInfo(Version, Built)
	app.CreateApplicationContext()

	os.Exit(cmd.Run(os.Args[1:]))


}
