/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package main

import (
	"github.com/arpabet/templateserv/cmd"
	"github.com/arpabet/templateserv/pkg/app"
	"math/rand"
	"os"
	"time"
)

var (
	Version   string
	Built     string
)

func main() {

	app.SetAppInfo(Version, Built)

	rand.Seed(time.Now().UnixNano())
	os.Exit(cmd.Run(os.Args[1:]))

}
