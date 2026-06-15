/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package dnsmod

import (
	"go.arpabet.com/sprint/dnsmod/netlify"
)

var DNSServices = []interface{} {
	WhoisService(),
	netlify.NetlifyProvider(),
}

