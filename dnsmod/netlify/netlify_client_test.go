/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package netlify_test

import (
	"fmt"
	"go.arpabet.com/sprint/dns"
	"go.arpabet.com/sprint/dnsmod"
	"go.arpabet.com/sprint/dnsmod/netlify"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestGetZone(t *testing.T) {

	domain := "www.example.com."

	fqdn := dnsmod.ToFqdn(domain)
	zone := FindZoneByFqdn(fqdn)

	secondLevel := dnsmod.UnFqdn(zone)
	println(secondLevel)

}

func noTestNetlifyMxChange(t *testing.T) {

	domain := "example.com"

	token := os.Getenv("NETLIFY_TOKEN")
	require.True(t, token != "")

	client := netlify.NewClient(token)

	publicIP, err := client.GetPublicIP()
	require.NoError(t, err)

	println(publicIP)

	fqdn := dnsmod.ToFqdn(domain)
	zone := FindZoneByFqdn(fqdn)

	fmt.Printf("zone=%v\n", zone)

	zone = dnsmod.UnFqdn(zone)

	list, err := client.GetRecords(zone)
	require.NoError(t, err)

	mxHostname := fmt.Sprintf("mx.%s", zone)

	createARecord := true
	createMXRecord := true
	for _, rec := range list {
		deleteRecord := false

		switch rec.Type {
		case "A":
			if strings.EqualFold(rec.Hostname, mxHostname) {
				if rec.Value == publicIP {
					createARecord = false
				} else {
					deleteRecord = true
				}
			}

		case "MX":
			if strings.EqualFold(rec.Value, mxHostname) {
				createMXRecord = false
			} else {
				deleteRecord = true
			}
		}

		if deleteRecord {
			fmt.Printf("DeleteRecord %v\n", rec)
			err = client.RemoveRecord(zone, rec.ID)
			require.NoError(t, err)
		}
	}

	if createARecord {

		record := &dns.DNSRecord{
			Hostname: "mx",
			TTL:      300,
			Type:     "A",
			Value:    publicIP,
		}

		record, err = client.CreateRecord(zone, record)
		require.NoError(t, err)

		fmt.Printf("Created Record %v\n", record)
	}

	if createMXRecord {

		record := &dns.DNSRecord{
			Hostname: zone,
			TTL:      300,
			Type:     "MX",
			Priority:  10,
			Value:    mxHostname,
		}

		record, err = client.CreateRecord(zone, record)
		require.NoError(t, err)

		fmt.Printf("Created Record %v\n", record)
	}


}

func FindZoneByFqdn(fqdn string) string {
	parts := strings.Split(fqdn, ".")
	switch len(parts) {
	case 1:
		return fqdn
	case 2:
		return fqdn
	default:
		return strings.Join(parts[len(parts)-2:], ".")
	}
}