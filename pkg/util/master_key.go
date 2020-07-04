/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package util

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/arpabet/template-server/pkg/constants"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"io"
)

func GenerateMasterKey() (string, error) {
	nonce := make([]byte, constants.KeySize)
	if _, err := io.ReadFull(rand.Reader, nonce); err == nil {
		key := constants.Encoding.EncodeToString(nonce)
		return key, nil
	} else {
		return "", err
	}
}

func ParseMasterKey(base64key string) ([]byte, error) {
	key, err := constants.Encoding.DecodeString(base64key)
	if err != nil {
		return key, err
	}
	if len(key) != constants.KeySize {
		return key, errors.Errorf("wrong key size %d", len(key))
	}
	return key, nil
}

func GetKeyHash(base64key string) string {
	hash := sha3.New256()
	hash.Write([]byte(base64key))
	digest := hex.EncodeToString(hash.Sum(nil))
	return digest[:8]
}