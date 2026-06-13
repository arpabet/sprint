/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package netlify

import (
	"go.arpabet.com/glue"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/netlify"
	"github.com/pkg/errors"
	"go.arpabet.com/sprint/cert"
	"os"
)

type implNetlifyChallenge struct {
	Properties   glue.Properties  `inject`
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
		return errors.Errorf("expected *lego.Client instance")
	}

	if token == "" {
		token = t.Properties.GetString("netlify.token", "")
	}

	if token == "" {
		token = os.Getenv("NETLIFY_TOKEN")
	}

	if token == "" {
		return errors.New("netlify token not found")
	}

	conf := netlify.NewDefaultConfig()
	conf.Token = token

	prov, err := netlify.NewDNSProviderConfig(conf)
	if err != nil {
		return err
	}

	return client.Challenge.SetDNS01Provider(prov)
}

