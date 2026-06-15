/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sealmod_test

import (
	"go.arpabet.com/sprint/sealmod"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRSASigner(t *testing.T) {

	text := "Hello World"
	plaintext := []byte(text)

	ss := sealmod.SealService()
	alice, err := ss.IssueSigner("RSA", 2048)
	require.NoError(t, err)

	signature, err := alice.Sign(plaintext)
	require.NoError(t, err)

	// RSA do not need sender public key
	actual, err := alice.Verify(plaintext, signature)
	require.NoError(t, err)

	require.True(t, actual)
}
