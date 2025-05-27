package limiter

import "time"

type LeakyBucket struct {
	capacity int
	rate     time.Duration
	queue    chan struct{}
}

func NewLeakyBucket(capacity int, ratePerSecond int) *LeakyBucket {
	lb := &LeakyBucket{
		capacity: capacity,
		rate:     time.Second / time.Duration(ratePerSecond),
		queue:    make(chan struct{}, capacity),
	}

	go lb.drain()
	return lb
}
func (lb *LeakyBucket) Allow() bool {
	select {
	// If enqueue succeeds, request is allowed
	case lb.queue <- struct{}{}:
		return true
	// If queue is full, reject the request
	default:
		return false
	}
}

func (lb *LeakyBucket) drain() {
	for {
		time.Sleep(lb.rate)
		select {

		case <-lb.queue:
			// Leak one request out of the bucket — this simulates processing it.

		default:
			// Nothing to leak, bucket is empty — do nothing.
		}
	}
}
