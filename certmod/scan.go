/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
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


