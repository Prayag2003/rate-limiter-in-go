package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity         int
	tokens           int
	refillRate       int
	lastRefilledTime time.Time
	mu               sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:         capacity,
		tokens:           capacity,
		refillRate:       refillRate,
		lastRefilledTime: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()

	// calculate how many tokens to refill
	elaspedTime := now.Sub(tb.lastRefilledTime).Seconds()
	tb.tokens += int(elaspedTime * float64(tb.refillRate))

	// don't exceed bucket capacity
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	// update last refill time
	tb.lastRefilledTime = now

	// allow request only if tokens are available
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}
