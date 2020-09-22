/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	"flag"
	"path/filepath"
)

const DAEMON_FLAG_KEY = "d"

var (
	address = flag.String("a", DefaultTlsAddress, "Address host:port")
	data = flag.String("data", "data", "Database Location Folder")
	log = flag.String("log", ExecutableLog, "Log File")
	useMmap = flag.Bool("mmap", false, "Use Memory Map Files")
	daemon = flag.Bool(DAEMON_FLAG_KEY, false, "Run as Daemon")
)

func GetArgs() []string {
	args := []string {
		"-a", *address,
		"-data", *data,
		"-log", *log,
	}
	if *useMmap {
		args = append(args, "-mmap")
	}
	return args
}


func ParseFlags(args []string) {
	flag.CommandLine.Parse(args)
}

func GetAddress() string {
	return *address
}

func GetLogFile() string {
	path, _ := filepath.Abs(*log)
	return path
}

func GetDataFolder() string {
	path, _ := filepath.Abs(*data)
	return path
}

func UseMemoryMap() bool {
	return *useMmap
}

func IsDaemon() bool {
	return *daemon
}

