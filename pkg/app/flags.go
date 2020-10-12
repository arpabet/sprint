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
	node = flag.String("node", DefaultNodeAddress, "Node Address host:port")
	data = flag.String("data", "data", "Database Location Folder")
	log = flag.String("log", "", "Log File")
	useMmap = flag.Bool("mmap", false, "Use Memory Map Files")
	daemon = flag.Bool(DAEMON_FLAG_KEY, false, "Run as Daemon")
)

func GetArgs() []string {
	args := []string {
		"-node", *node,
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

func GetNodeAddress() string {
	return *node
}

func GetDataFolder() string {
	path, _ := filepath.Abs(*data)
	return path
}

func GetLogFile() string {
	dir := *log
	if dir == "" {
		dir = ExecutableName + ".log"
	}
	path, _ := filepath.Abs(dir)
	return path
}

func UseMemoryMap() bool {
	return *useMmap
}

func IsDaemon() bool {
	return *daemon
}

