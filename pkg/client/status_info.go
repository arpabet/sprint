/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package client

import (
	"encoding/json"
	"fmt"
)

type StatusInfo struct {
	Build    string
	Version  string
	Started  string
	MasterKeyHash  string
}

func (t *StatusInfo) String() string {
	return fmt.Sprintf("MasterKeyHash %s, Build %s, Version %s, Started at %s", t.MasterKeyHash, t.Build, t.Version, t.Started)
}

func ParseStatusInfo(content []byte) (*StatusInfo, error) {
	si := &StatusInfo{}
	if err := json.Unmarshal(content, si); err == nil {
		return si, nil
	} else {
		return nil, err
	}
}

