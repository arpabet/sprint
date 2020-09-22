/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	"github.com/arpabet/templateserv/pkg/resources"
	"strings"
)

func GetLicenses() string {
	if content, err := resources.Asset("licenses.txt"); err == nil {
		return filterLines(string(content), ApplicationName)
	}
	return ""
}

func GetSwagger() string {
	if content, err := resources.Asset("swagger/server.swagger.json"); err == nil {
		return string(content)
	}
	return ""
}

func filterLines(content string, words ...string) string {

	var out strings.Builder

	for _, line := range strings.Split(content, "\n") {
		include := true
		for _, word := range words {
			if strings.Contains(line, word) {
				include = false
				break
			}
		}
		if include {
			out.WriteString(line)
			out.WriteRune('\n')
		}
	}

	return out.String()
}

