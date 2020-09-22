/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package client

import (
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/util"
	"io/ioutil"
	"net/url"
	"strings"
)

func RequestStatus() (string, error) {
	client, err := util.NewClient()
	if err != nil {
		return "", err
	}

	resp, err := client.Get(app.StatusEndpoint())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func RequestStop() (string, error) {
	client, err := util.NewClient()
	if err != nil {
		return "", err
	}

	resp, err := client.Post(app.StopEndpoint(), app.PlainContentType, strings.NewReader(""))
	if err != nil {
		return "", err
	}

	return util.PostResp(resp)
}

func SetConfig(key, value string) (string, error) {

	client, err := util.NewClient()
	if err != nil {
		return "", err
	}

	formData := url.Values{
		"key" : {key},
		"value": {value},
	}

	resp, err := client.PostForm(app.SetConfigEndpoint(), formData)
	if err != nil {
		return "", err
	}

	return util.PostResp(resp)
}

func GetConfig(key string) (string, error) {

	client, err := util.NewClient()
	if err != nil {
		return "", err
	}

	value, err := util.ReadAll(client.Get(app.GetConfigEndpoint(key)))
	if err != nil {
		return "", err
	}

	return value, nil
}

