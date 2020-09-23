/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package util

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/arpabet/timeuuid"
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"io"
	"strconv"
	"strings"
)

func GenerateMasterKey() (string, error) {
	nonce := make([]byte, app.KeySize)
	if _, err := io.ReadFull(rand.Reader, nonce); err == nil {
		key := app.Encoding.EncodeToString(nonce)
		return key, nil
	} else {
		return "", err
	}
}

func ParseMasterKey(base64key string) ([]byte, error) {
	key, err := app.Encoding.DecodeString(base64key)
	if err != nil {
		return key, err
	}
	if len(key) != app.KeySize {
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

func GenerateNodeId() (string, error) {

	blob := make([]byte, app.NodeIdSize)
	if _, err := io.ReadFull(rand.Reader, blob); err == nil {
		return "0x" + hex.EncodeToString(blob), nil
	} else {
		return "", err
	}

}

func ParseNodeId(nodeId string) (uint64, error) {
	if strings.HasPrefix(nodeId, "0x") {
		nodeId = nodeId[2:]
	}
	return strconv.ParseUint(nodeId, 16, app.NodeIdBits)
}

func EventKey(uuid timeuuid.UUID) ([]byte, error) {
	key := make([]byte, 1 + 16)
	key[0] = app.EventPrefix
	if err := uuid.MarshalSortableBinaryTo(key[1:]); err != nil {
		return nil, err
	}
	return key, nil
}
