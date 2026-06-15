/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package raftcmd

var RaftCommands = []interface{}{
	SerfJoinCommand(),
	SerfMembersCommand(),
	SerfEventCommand(),
	SerfInfoCommand(),
	SerfVersionCommand(),
	SerfLeaveCommand(),
	SerfMonitorCommand(),
	SerfReachabilityCommand(),
	SerfRttCommand(),
	SerfTagsCommand(),
	SerfCommands(),
}
