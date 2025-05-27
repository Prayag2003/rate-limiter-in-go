package main

import (
	"flag"
	"fmt"

	"github.com/Prayag2003/rate-limiter-in-go/config"
	"github.com/Prayag2003/rate-limiter-in-go/internal/limiter"
	"github.com/Prayag2003/rate-limiter-in-go/internal/simulator"
)

func main() {
	limiterTypeFlag := flag.String("type", "", "Rate limiter type override (token or leaky)")
	flag.Parse()

	cfg := config.LoadConfig()

	// Override from CLI if provided
	if *limiterTypeFlag != "" {
		fmt.Println("[CLI] Overriding limiter type to:", *limiterTypeFlag)
		cfg.RateLimiterType = *limiterTypeFlag
	}

	// Choose limiter
	var limiterInstance limiter.RateLimiter
	switch cfg.RateLimiterType {
	case "token":
		limiterInstance = limiter.NewTokenBucket(cfg.Capacity, cfg.RefillRate)
	case "leaky":
		limiterInstance = limiter.NewLeakyBucket(cfg.Capacity, cfg.LeakRate)
	default:
		panic("Unsupported RATE_LIMITER_TYPE: " + cfg.RateLimiterType)
	}

	simulator.RunRealisticSimulation(
		limiterInstance,
		cfg.RPS,
		cfg.DurationSec,
		cfg.Concurrency,
	)
}
