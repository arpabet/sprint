/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintutils

import (
	"crypto/tls"
)

func ParseClientAuth(s string) tls.ClientAuthType {
	switch s {
	case "no_client_cert":
		return tls.NoClientCert
	case "request_client_cert":
		return tls.RequestClientCert
	case "require_any_client_cert":
		return tls.RequireAnyClientCert
	case "verify_client_cert":
		return tls.VerifyClientCertIfGiven
	case "require_verify_client_cert":
		return tls.RequireAndVerifyClientCert
	default:
		return tls.NoClientCert
	}
}
