/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprint


const (

	/**
	Predefined role name for the core context, implementing base functionality of the application.
	Sprint Framework always tries to create core context first, and then extend it to servers.
	*/
	CoreRole = "core"

	/**
	Predefined role name for the server context in the application.
	Usually can container multiple child contexts for each server.
	*/
	ServerRole = "server"

	/**
	Predefined role name for the client connecting to the node by using control RPC.
	*/
	ControlClientRole = "control-client"


)