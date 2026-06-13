/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
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
