package simulator

import (
	"fmt"
	"sync"
	"time"

	"github.com/Prayag2003/rate-limiter-in-go/internal/limiter"
)

func RunRealisticSimulation(limiter limiter.RateLimiter, rps, durationSec, concurrency int) {
	totalRequests := rps * durationSec
	var wg sync.WaitGroup
	var mu sync.Mutex

	allowed := 0
	rejected := 0

	requests := make(chan struct{}, totalRequests)
	for i := 0; i < totalRequests; i++ {
		requests <- struct{}{}
	}
	close(requests)

	start := time.Now()

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			ticker := time.NewTicker(time.Second / time.Duration(rps/concurrency))
			defer ticker.Stop()

			for range ticker.C {
				_, ok := <-requests
				if !ok {
					return
				}

				if limiter.Allow() {
					mu.Lock()
					allowed++
					mu.Unlock()
				} else {
					mu.Lock()
					rejected++
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("[Simulation] RPS: %d | Duration: %ds | Allowed: %d | Rejected: %d | Time: %.2fs\n",
		rps, durationSec, allowed, rejected, elapsed.Seconds())
}
