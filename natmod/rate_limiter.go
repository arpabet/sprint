/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package natmod

import (
	"sync"
	"time"
)

const (
	defaultRateLimit  = 100 * time.Millisecond
)

type RateLimiter struct {
	Limit  time.Duration
	mu  sync.Mutex
	lastReqTime time.Time
}

func (t *RateLimiter) Do(fn func() error) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Limit == 0 {
		t.Limit = defaultRateLimit
	}

	lastreq := time.Since(t.lastReqTime)
	if lastreq < t.Limit {
		time.Sleep(t.Limit - lastreq)
	}
	err := fn()
	t.lastReqTime = time.Now()
	return err
}

