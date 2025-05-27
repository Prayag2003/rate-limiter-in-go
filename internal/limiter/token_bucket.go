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
