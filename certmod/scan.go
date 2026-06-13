/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package certmod

import (
	"go.arpabet.com/sprint/certmod/netlify"
)

var CertServices = []interface{} {
	CertificateIssueService(),
	CertificateRepository(),
	CertificateService(),
	CertificateManager(),
	netlify.NetlifyChallenge(),
	DynDNSService(),
}


