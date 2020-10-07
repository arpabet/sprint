/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package util

import (
	"github.com/arpabet/sprint/pkg/resources"
	"text/template"
)

func MustAssetTemplate(name string) *template.Template {
	asset := resources.MustAsset(name)
	return template.Must(template.New(name).Parse(string(asset)))
}

