/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package util

import "time"

func Hours(d time.Duration) float64 {
	sec := float64(d) / float64(time.Hour)
	return sec
}

func Seconds(d time.Duration) float64 {
	sec := float64(d) / float64(time.Second)
	return sec
}

