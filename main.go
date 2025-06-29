package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Prayag2003/rate-limiter-in-go/config"
	"github.com/Prayag2003/rate-limiter-in-go/internal/limiter"
	"github.com/Prayag2003/rate-limiter-in-go/middleware"
)

func main() {
	limiterTypeFlag := flag.String("type", "", "Rate limiter type override (token or leaky)")
	flag.Parse()

	cfg := config.LoadConfig()

	if *limiterTypeFlag != "" {
		fmt.Println("[CLI] Overriding limiter type to:", *limiterTypeFlag)
		cfg.RateLimiterType = *limiterTypeFlag
	}

	var limiterInstance limiter.RateLimiter
	switch cfg.RateLimiterType {
	case "token":
		limiterInstance = limiter.NewTokenBucket(cfg.Capacity, cfg.RefillRate)
	case "leaky":
		limiterInstance = limiter.NewLeakyBucket(cfg.Capacity, cfg.LeakRate)
	default:
		panic("Unsupported RATE_LIMITER_TYPE: " + cfg.RateLimiterType)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	wrappedMux := middleware.RateLimiterMiddleware(limiterInstance)(mux)

	fmt.Println("Starting HTTP server on :8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
