package simulator

import (
	"fmt"
	"sync"
	"time"

	"github.com/Prayag2003/rate-limiter-in-go/internal/limiter"
)

func SimulateLimiter(name string, rl limiter.RateLimiter, totalRequests int, concurrent int) {
	var wg sync.WaitGroup
	var allowed, rejected int64

	start := time.Now()

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < totalRequests/concurrent; j++ {
				if rl.Allow() {
					allowed++
				} else {
					rejected++
				}
			}
			time.Sleep(5 * time.Second)
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("[%s] Allowed: %d | Rejected: %d | Simulation Duration: %s\n", name, allowed, rejected, duration)
}
