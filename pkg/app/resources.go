/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	"io/ioutil"
	"strings"
)

func GetLicenses() string {
	if file, err := Resources.Open("licenses.txt"); err == nil {
		if content, err := ioutil.ReadAll(file); err == nil {
			return filterLines(string(content), PackageName)
		}
	}
	return ""
}

func GetSwagger() string {
	if file, err := Resources.Open("swagger/node.swagger.json"); err == nil {
		if content, err := ioutil.ReadAll(file); err == nil {
			return string(content)
		}
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

