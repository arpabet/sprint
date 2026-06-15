/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintcore

var CoreServices = []interface{} {
	ZapLogFactory(),
	HCLogFactory(),
	NodeService(),
	ConfigRepository(10000),
	JobService(),
	StorageService(),
	MailService(),
}
