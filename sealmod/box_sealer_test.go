/*
 * Copyright (c) 2025 Karagatan LLC.
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

func TestBOXSealer(t *testing.T) {

	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	require.NoError(t, err)

	text := "Hello World"
	plaintext := []byte(text)

	ss := sealmod.SealService()
	// box do not need key length, always 256 bit
	alice, err := ss.IssueSealer("BOX", 0)
	require.NoError(t, err)

	bob, err := ss.IssueSealer("BOX", 0)
	require.NoError(t, err)

	ciphertext, err := alice.Seal(plaintext, bob.PublicKey())
	require.NoError(t, err)

	actual, err := bob.Open(ciphertext, alice.PublicKey())
	require.NoError(t, err)

	require.True(t, bytes.Equal(plaintext, actual))
}
