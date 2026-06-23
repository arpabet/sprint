/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintutils

import (
	"runtime/debug"

	"golang.org/x/xerrors"
)

func PanicToError(err *error) {
	if r := recover(); r != nil {
		*err = xerrors.Errorf("%v, %s", r, debug.Stack())
	}
}
