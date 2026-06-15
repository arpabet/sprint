/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package natmod

import (
	"errors"
	"go.arpabet.com/sprint/nat"
	"net"
	"time"
)

var (
	ErrNoNatService = errors.New("no nat service")
)

type implNonatService struct {
}

func NoNatService() nat.NatService {
	return &implNonatService{}
}

func (t *implNonatService) ServiceName() string {
	return "no_nat"
}

func (t *implNonatService) AllowMapping() bool {
	return false
}

func (t *implNonatService) AddMapping(protocol string, extport, intport int, name string, lifetime time.Duration) error {
	return nil
}

func (t *implNonatService) DeleteMapping(protocol string, extport, intport int) error {
	return nil
}

func (t *implNonatService) ExternalIP() (net.IP, error) {
	return nil, ErrNoNatService
}


