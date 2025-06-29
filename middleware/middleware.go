package middleware

import (
	"net/http"

	"github.com/Prayag2003/rate-limiter-in-go/internal/limiter"
)

func RateLimiterMiddleware(l limiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !l.Allow() {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("Rate limit exceeded"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
