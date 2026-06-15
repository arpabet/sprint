/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintutils_test

import (
	"go.arpabet.com/sprint/sprintframework/sprintutils"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestParseFileMode(t *testing.T) {

	knownModes := map[string]os.FileMode{
		"-rwxrwxr-x":    os.FileMode(0775),
		"-rw-rw-r--":    os.FileMode(0664),
		"-rw-rw-rw-":    os.FileMode(0666),
		"-rwxrwx---":    os.FileMode(0770),
	}

	for expected, mode := range knownModes {

		str := mode.String()
		require.Equal(t, expected, str)

		actual := sprintutils.ParseFileMode(str)
		require.Equal(t, mode, actual, mode.String())

	}

}
