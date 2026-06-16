package ratelimit

import (
	"sync/atomic"
	"time"
)

type AtomicRateLimiter struct {
	tokens int64
	max    int64
}

func NewAtomicRateLimiter(max int64) *AtomicRateLimiter {
	if max < 0 {
		max = 0
	}
	tokens := max
	return &AtomicRateLimiter{tokens: tokens, max: max}
}

func (r *AtomicRateLimiter) Allow() bool {
	for {
		cur := atomic.LoadInt64(&r.tokens)
		if cur <= 0 {
			return false
		}
		if atomic.CompareAndSwapInt64(&r.tokens, cur, cur-1) {
			return true
		}
	}
}

func (r *AtomicRateLimiter) refill(n int64) {
	for {
		cur := atomic.LoadInt64(&r.tokens)
		next := min(cur+n, r.max)
		if atomic.CompareAndSwapInt64(&r.tokens, cur, next) {
			return
		}
	}
}

func (r *AtomicRateLimiter) Start(interval time.Duration, amount int64) func() {
	if interval <= 0 || amount <= 0 {
		return func() {}
	}
	stopCh := make(chan struct{})

	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				r.refill(amount)
			case <-stopCh:
				return
			}
		}
	}()

	return func() { close(stopCh) }
}
