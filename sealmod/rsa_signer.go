/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */
package sealmod

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"reflect"

	"go.arpabet.com/sprint/seal"
	"golang.org/x/xerrors"
)

type implRSASigner struct {
	pub  *rsa.PublicKey
	priv *rsa.PrivateKey
}

func RSASigner(opt *seal.SealerOptions) (seal.AsymmetricSigner, error) {
	t := &implRSASigner{}

	if opt.PublicKey != nil {
		var ok bool
		t.pub, ok = opt.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, xerrors.Errorf("not a RSA public key, %v", reflect.TypeOf(opt.PublicKey))
		}
	}
	if opt.PrivateKey != nil {
		var ok bool
		t.priv, ok = opt.PrivateKey.(*rsa.PrivateKey)
		if !ok {
			return nil, xerrors.Errorf("not a RSA private key, %v", reflect.TypeOf(opt.PrivateKey))
		}
	}

	return t, nil
}

func RSASignerIssue(bits int) (seal.AsymmetricSigner, error) {
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return &implRSASigner{
		pub:  &priv.PublicKey,
		priv: priv,
	}, nil
}

func (t *implRSASigner) PublicKey() crypto.PublicKey {
	return t.pub
}

func (t *implRSASigner) PrivateKey() crypto.PrivateKey {
	return t.priv
}

func (t *implRSASigner) EncodePublicKey() (string, error) {

	if t.pub == nil {
		return "", xerrors.New("public key is empty")
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(t.pub)
	if err != nil {
		return "", xerrors.Errorf("marshal PKIX public key, %v", err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return string(pubBytes), nil
}

func (t *implRSASigner) EncodePrivateKey() (string, error) {
	if t.priv == nil {
		return "", xerrors.New("private key is empty")
	}
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(t.priv),
		},
	)
	return string(privBytes), nil
}

func (t *implRSASigner) Sign(plaintext []byte) (signature []byte, err error) {
	if t.priv == nil {
		return nil, xerrors.New("private key is empty")
	}
	hash := sha512.New()
	hash.Write(plaintext)
	signature, err = rsa.SignPKCS1v15(rand.Reader, t.priv, crypto.SHA512, hash.Sum(nil))
	return
}

func (t *implRSASigner) Verify(plaintext, signature []byte) (valid bool, err error) {
	if t.pub == nil {
		return false, xerrors.New("public key is empty")
	}
	hash := sha512.New()
	hash.Write(plaintext)
	err = rsa.VerifyPKCS1v15(t.pub, crypto.SHA512, hash.Sum(nil), signature)
	return err == nil, err
}
