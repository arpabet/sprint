/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package natmod

import (
	"github.com/pkg/errors"
	"go.arpabet.com/sprint/nat"
	"net"
	"time"
)

type implExternalIPService struct {
	ip  net.IP
}

func ExternalIPService(address string) (nat.NatService, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return nil, errors.Errorf("invalid IP address '%s'", address)
	}
	return &implExternalIPService{ip: ip}, nil
}

func (t *implExternalIPService) ServiceName() string {
	return "ext_ip"
}

func (t *implExternalIPService) AllowMapping() bool {
	return false
}

func (t *implExternalIPService) AddMapping(protocol string, extport, intport int, name string, lifetime time.Duration) error {
	return nil
}

func (t *implExternalIPService) DeleteMapping(protocol string, extport, intport int) error {
	return nil
}

func (t *implExternalIPService) ExternalIP() (net.IP, error) {
	return t.ip, nil
}
