/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package app

import "fmt"

func StatusEndpoint() string {
	return fmt.Sprintf(statusEndpoint, GetAddress())
}

func StopEndpoint() string {
	return fmt.Sprintf(stopEndpoint, GetAddress())
}

func SetConfigEndpoint() string {
	return fmt.Sprintf(setConfigEndpoint, GetAddress())
}

func GetConfigEndpoint(key string) string {
	return fmt.Sprintf(getConfigEndpoint, GetAddress(), key)
}
