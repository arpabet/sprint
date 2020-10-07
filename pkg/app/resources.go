/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package app

import (
	"io/ioutil"
	"strings"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

func Asset(name string) ([]byte, error) {
	asset, err := Assets.Open(name)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(asset)
}

func Resource(name string) ([]byte, error) {
	asset, err := Resources.Open(name)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(asset)
}

func ResourceTextTemplate(name string) (*textTemplate.Template, error) {
	content, err := Resource(name)
	if err != nil {
		return nil, err
	}
	return textTemplate.New(name).Parse(string(content))
}

func ResourceHtmlTemplate(name string) (*htmlTemplate.Template, error) {
	content, err := Resource(name)
	if err != nil {
		return nil, err
	}
	return htmlTemplate.New(name).Parse(string(content))
}

func GetLicenses() string {
	if content, err := Resource(LicensesFile); err == nil {
		return filterLines(string(content), PackageName)
	}
	return ""
}

func GetSwagger() string {
	if content, err := Resource(SwaggerFile); err == nil {
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

