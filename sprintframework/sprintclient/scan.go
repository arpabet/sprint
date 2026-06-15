/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintclient

var ControlClientBeans = []interface{} {
	GrpcClientFactory("control-grpc-client"),
	ControlClient(),
}

