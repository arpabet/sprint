/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package netlify

import (
	"os"

	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/netlify"
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/cert"
	"golang.org/x/xerrors"
)

type implNetlifyChallenge struct {
	Properties glue.Properties `inject:""`
}

func NetlifyChallenge() cert.DNSChallenge {
	return &implNetlifyChallenge{}
}

func (t *implNetlifyChallenge) BeanName() string {
	return "netlify_challenge"
}

func (t *implNetlifyChallenge) RegisterChallenge(legoClient interface{}, token string) error {

	client, ok := legoClient.(*lego.Client)
	if !ok {
		return xerrors.Errorf("expected *lego.Client instance")
	}

	if token == "" {
		token = t.Properties.GetString("netlify.token", "")
	}

	if token == "" {
		token = os.Getenv("NETLIFY_TOKEN")
	}

	if token == "" {
		return xerrors.New("netlify token not found")
	}

	conf := netlify.NewDefaultConfig()
	conf.Token = token

	prov, err := netlify.NewDNSProviderConfig(conf)
	if err != nil {
		return err
	}

	return client.Challenge.SetDNS01Provider(prov)
}
