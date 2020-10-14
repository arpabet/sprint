/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	"flag"
	"path/filepath"
)

const (

	DAEMON_FLAG_KEY = "d"
	DISTR_FLAG_KEY = "distr"
)

var (
	node = flag.String("node", "", "Node Address host:port")
	data = flag.String("data", "", "Database Location Folder")
	log = flag.String("log", "", "Log File")
	distr = flag.String(DISTR_FLAG_KEY, "", "Distr File")
	useMmap = flag.Bool("mmap", false, "Use Memory Map Files")
	daemon = flag.Bool(DAEMON_FLAG_KEY, false, "Run as Daemon")
)

func GetArgs() []string {
	var args []string
	if *node != "" {
		args = append(args, "-node", *node)
	}
	if *data != "" {
		args = append(args, "-data", *data)
	}
	if *log != "" {
		args = append(args, "-log", *log)
	}
	if *useMmap {
		args = append(args, "-mmap")
	}
	return args
}

func ParseFlags(args []string) {
	flag.CommandLine.Parse(args)
}

func GetNodeAddress() string {
	value := *node
	if value == "" {
		value = DefaultNodeAddress
	}
	return value
}

func GetDistrFile() string {
	return *distr
}

func GetDataFolder() string {
	value := *data
	if value == "" {
		return ExecutableData()
	} else {
		path, _ := filepath.Abs(value)
		return path
	}
}

func GetLogFile() string {
	value := *log
	if value == "" {
		return ExecutableLog()
	} else {
		path, _ := filepath.Abs(value)
		return path
	}
}

func UseMemoryMap() bool {
	return *useMmap
}

func IsDaemon() bool {
	return *daemon
}

