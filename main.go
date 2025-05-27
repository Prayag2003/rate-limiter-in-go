package main

import (
	"github.com/Prayag2003/rate-limiter-in-go/internal/limiter"
	"github.com/Prayag2003/rate-limiter-in-go/internal/simulator"
)

func main() {
	tokenBucketLimiter := limiter.NewTokenBucket(100, 5)
	leakyBucketLimiter := limiter.NewLeakyBucket(100, 50)

	simulator.SimulateLimiter("TokenBucket", tokenBucketLimiter, 1000, 10)
	simulator.SimulateLimiter("LeakyBucket", leakyBucketLimiter, 1000, 10)
}
