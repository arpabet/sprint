/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintutils

import (
	"github.com/pkg/errors"
	"runtime/debug"
)

func PanicToError(err *error) {
	if r := recover(); r != nil {
		*err = errors.Errorf("%v, %s", r, debug.Stack())
	}
}

