/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sealmod_test

import (
	"bytes"
	"crypto/rand"
	"go.arpabet.com/sprint/sealmod"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestRSASealer(t *testing.T) {

	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	require.NoError(t, err)

	text := "Hello World"
	plaintext := []byte(text)

	ss := sealmod.SealService()
	alice, err := ss.IssueSealer("RSA", 2048)
	require.NoError(t, err)

	ciphertext, err := alice.Seal(plaintext, alice.PublicKey())
	require.NoError(t, err)

	// RSA do not need sender public key
	actual, err := alice.Open(ciphertext, nil)
	require.NoError(t, err)

	require.True(t, bytes.Equal(plaintext, actual))
}
