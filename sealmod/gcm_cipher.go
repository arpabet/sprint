/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sealmod

import (
	"crypto/cipher"
	"crypto/rand"
	"go.arpabet.com/sprint/seal"
	"github.com/pkg/errors"
	"io"
)

type implGCMCipher struct {
	gcm  cipher.AEAD
}

func GCMCipher(opt *seal.CipherOptions) (this seal.AuthenticatedCipher, err error) {
	t := &implGCMCipher{}
	t.gcm, err = cipher.NewGCM(opt.Block)
	return t, err
}

func (t *implGCMCipher) Key() cipher.AEAD {
	return t.gcm
}

func (t *implGCMCipher) Encrypt(plaintext []byte) (ciphertext []byte, err error) {
	nonce := make([]byte, t.gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return t.gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (t *implGCMCipher) Decrypt(ciphertext []byte) (plaintext []byte, err error) {

	nonceSize := t.gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.Errorf("ciphertext len %d is less than GCM nonce size %d", len(ciphertext), nonceSize)
	}

	nonce, encrypted := ciphertext[:nonceSize], ciphertext[nonceSize:]

	return t.gcm.Open(nil, nonce, encrypted, nil)
}