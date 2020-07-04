/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package client

import (
	"github.com/arpabet/template-server/pkg/constants"
	"github.com/arpabet/template-server/pkg/util"
	"io/ioutil"
	"fmt"
)

func RequestStatus() (*StatusInfo, error) {
	client, err := util.NewClient()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(constants.StatusEndpoint())

	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ParseStatusInfo(body)
}

