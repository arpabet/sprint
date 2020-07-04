/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package constants

import (
	"flag"
	"path/filepath"
	"fmt"
)

var (
	address = flag.String("addr", "localhost:8080", "Address host:port")
	database = flag.String("data", "data", "Database Location Folder")
	ssl = flag.String("ssl", "ssl", "SSL Certificates Folder")
	useMmap = flag.Bool("mmap", false, "Use Memory Map Files")
	debug = flag.Bool("debug", false, "Use Debug Log Mode")
	daemon = flag.Bool("d", false, "Run as Daemon")
)

const DAEMON_FLAG_KEY = "d"

func GetArgs() []string {
	args := []string {
		"-addr", *address,
		"-data", *database,
	}
	if *useMmap {
		args = append(args, "-mmap")
	}
	if *debug {
		args = append(args, "-debug")
	}
	return args
}


func ParseFlags() {
	flag.Parse()
}

func GetAddress() string {
	return *address
}

func GetSSLFolder() string {
	path, _ := filepath.Abs(*ssl)
	return path
}

func GetDatabaseFolder() string {
	path, _ := filepath.Abs(*database)
	return path
}

func UseMemoryMap() bool {
	return *useMmap
}

func IsDebug() bool {
	return *debug
}

func IsDaemon() bool {
	return *daemon
}

func StatusEndpoint() string {
	return fmt.Sprintf(statusEndpoint, GetAddress())
}