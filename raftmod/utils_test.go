/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package raftmod_test

import (
	"fmt"
	"go.arpabet.com/sprint/raftmod"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLocalIP(t *testing.T) {

	ip, err := raftmod.PrivateIP()
	require.NoError(t, err)

	fmt.Printf("ip=%v\n", ip)


}
