/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package fsmod

var EmptyValues = map[string]bool {
	"n/a": true,
	"N/A": true,
	"N/a": true,
	"null": true,
	"NULL": true,
	"Null": true,
	"nil": true,
	"NIL": true,
	"Nil": true,
	"nan": true,
	"NaN": true,
	"Nan": true,
	"#": true,
	"": true,
}

func PandasFriendly(v string) string {
	if EmptyValues[v] {
		return "#"
	} else {
		return v
	}
}

func RemoveHash(v string) string {
	if v == "#" {
		return ""
	}
	return v
}

