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

func (t *configService) GetAll(cb func(key, value string)) error {
	prefix := []byte(app.ConfigPrefix)
	return t.Storage.Enumerate(prefix, prefix, 256, func(key, value []byte) bool {
		configKey := key[app.ConfigPrefixLen:]
		cb(string(configKey), string(value))
		return true
	})
}

func (t *configService) GetWithDefault(key, defaultValue string) (string, error) {
	value, err := t.Storage.Get(t.toBin(key), false)
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

func (t *configService) GetInt(key string, defaultValue int) (int, error) {
	str, err := t.Get(key)
	if err != nil {
		return defaultValue, err
	}
	if str == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(str)
}

func (t *configService) Set(key, value string) error {
	if value == "" {
		return t.Storage.Remove(t.toBin(key))
	} else {
		return t.Storage.Put(t.toBin(key), []byte(value))
	}
}

func (t *configService) toBin(key string) []byte {
	return []byte(app.ConfigPrefix + key)
}

