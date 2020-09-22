/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package service

import (
	"github.com/arpabet/templateserv/pkg/app"
	"strconv"
)

type configService struct {
	app.Storage  `inject`
}

func ConfigService() app.ConfigService {
	return &configService{}
}

func (t *configService) Get(key string) (string, error) {
	return t.GetWithDefault(key, "")
}

func (t *configService) GetWithDefault(key, defaultValue string) (string, error) {
	value, err := t.Storage.Get([]byte(app.ConfigPrefix + key), false)
	if err != nil {
		return "", err
	} else if value != nil {
		return string(value), nil
	} else {
		return defaultValue, nil
	}
}

func (t *configService) GetBool(key string) (bool, error) {
	str, err := t.Get(key)
	if err != nil {
		return false, err
	}
	if str == "" {
		return false, nil
	}
	return strconv.ParseBool(str)
}

func (t *configService) Set(key, value string) error {
	if value == "" {
		return t.Storage.Remove([]byte(app.ConfigPrefix + key))
	} else {
		return t.Storage.Put([]byte(app.ConfigPrefix + key), []byte(value))
	}
}
