/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package main

import (
	"github.com/arpabet/templateserv/pkg/cmd"
	"github.com/arpabet/templateserv/pkg/app"
	"math/rand"
	"os"
	"time"
)

var (
	Version   string
	Build     string
)

func main() {

	app.Version = Version
	app.Build = Build

	rand.Seed(time.Now().UnixNano())
	os.Exit(cmd.Run(os.Args[1:]))

}
