/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package cert

import (
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/dns"
	"reflect"
)

var DNSChallengeClass = reflect.TypeOf((*DNSChallenge)(nil)).Elem()

type DNSChallenge interface {
	glue.NamedBean

	RegisterChallenge(legoClient interface{}, token string) error

}


var DynDNSServiceClass = reflect.TypeOf((*DynDNSService)(nil)).Elem()

type DynDNSService interface {
	glue.NamedBean
	glue.InitializingBean

	EnsureAllPublic(subDomains ...string) error

	EnsureCustom(func(client dns.DNSProviderClient, zone string, externalIP string) error) error

}
