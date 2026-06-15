/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintcmd

var ApplicationCommands = []interface{}{
	VersionCommand(),
	SetupCommand(),
	HelpCommand(),
	ResourcesCommand(),
	ConfigCommand(),
	CertsCommand(),
	StorageCommand(),
	JobsCommand(),
	KeygenCommand(),
	NodeCommand(),
	RunNode(),
	StartNode(),
	StopNode(),
	RestartNode(),
	StatusNode(),
}
