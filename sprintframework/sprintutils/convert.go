/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintutils

import "github.com/pkg/errors"

func ParseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "on", "ON", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "off", "OFF", "Off":
		return false, nil
	}
	return false, errors.Errorf("invalid syntax %s", str)
}

