/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package nat

import (
	"net"
	"reflect"
	"time"
)

var NatServiceClass = reflect.TypeOf((*NatService)(nil)).Elem()

type NatService interface {

	AllowMapping() bool

	AddMapping(protocol string, extport, intport int, name string, lifetime time.Duration) error

	DeleteMapping(protocol string, extport, intport int) error

	ExternalIP() (net.IP, error)

	ServiceName() string
}
