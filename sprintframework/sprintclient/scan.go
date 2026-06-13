/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package sprintclient

var ControlClientBeans = []interface{} {
	GrpcClientFactory("control-grpc-client"),
	ControlClient(),
}

