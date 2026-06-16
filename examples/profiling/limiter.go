package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu     sync.Mutex
	tokens int64
	max    int64
}

func NewRateLimiter(tokens, max int64) *RateLimiter {
	return &RateLimiter{tokens: tokens, max: max}
}
func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.tokens <= 0 {
		return false
	}
	r.tokens--
	return true
}
func (r *RateLimiter) refill(n int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tokens = min(r.tokens+n, r.max)
}
func (r *RateLimiter) StartRefill(interval time.Duration, amount int64) func() {
	if interval <= 0 || amount <= 0 {
		return func() {}
	}

	stopCh := make(chan struct{})

	r.refillTicker(stopCh, interval, amount)

	return func() { close(stopCh) }
}
func (r *RateLimiter) refillTicker(stopCh chan struct{}, interval time.Duration, amount int64) {

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				r.refill(amount)
			case <-stopCh:
				return
			}
		}
	}()
}
