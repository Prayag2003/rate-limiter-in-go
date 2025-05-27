package limiter

type LeakyBucket struct {
	capacity int
	rate     int
	queue    chan struct{}
}
