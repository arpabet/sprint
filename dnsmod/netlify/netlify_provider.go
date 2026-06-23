/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package netlify

import (
	"strings"

	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/dns"
	"golang.org/x/xerrors"
)

type implNetlifyProvider struct {
	Properties glue.Properties `inject:""`
}

func NetlifyProvider() dns.DNSProvider {
	return &implNetlifyProvider{}
}

func (t *implNetlifyProvider) BeanName() string {
	return "netlify_provider"
}

func (t *implNetlifyProvider) Detect(whois *dns.Whois) bool {
	for _, ns := range whois.NServer {
		if strings.HasSuffix(strings.ToLower(ns), ".nsone.net") {
			return true
		}
	}
	return false
}

func (t *implNetlifyProvider) NewClient(token string) (dns.DNSProviderClient, error) {

	/*
		if token == "" {
			token = t.Properties.GetString("netlify.token", "")
		}

		if token == "" {
			token = os.Getenv("NETLIFY_TOKEN")
		}
	*/

	if token == "" {
		return nil, xerrors.New("netlify token is empty")
	}

	return NewClient(token), nil
}
