/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

type AppInfo struct {
	Version   string
	Build     string
}

var appInfo AppInfo

func SetAppInfo(version, build string) {
	appInfo.Version = version
	appInfo.Build = build
}

func GetAppInfo() *AppInfo {
	return &appInfo
}
