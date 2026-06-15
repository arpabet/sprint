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

func TestGCMCipher(t *testing.T) {

	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	require.NoError(t, err)

	text := "Hello World"
	plaintext := []byte(text)

	ss := sealmod.SealService()
	gcm, err := ss.AuthenticatedCipher(sealmod.WithCipher("GCM"), sealmod.WithAESKey(key))
	require.NoError(t, err)

	ciphertext, err := gcm.Encrypt(plaintext)
	require.NoError(t, err)

	actual, err := gcm.Decrypt(ciphertext)
	require.NoError(t, err)

	require.True(t, bytes.Equal(plaintext, actual))
}

