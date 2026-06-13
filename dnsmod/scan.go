/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package dnsmod

import (
	"go.arpabet.com/sprint/dnsmod/netlify"
)

var DNSServices = []interface{} {
	WhoisService(),
	netlify.NetlifyProvider(),
}

