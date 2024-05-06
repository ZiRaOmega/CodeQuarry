package app

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu         sync.Mutex
	rateLimit  int
	timeWindow time.Duration
	counters   map[string]int
	timestamps map[string]time.Time
}

func NewRateLimiter(rateLimit int, timeWindow time.Duration) *RateLimiter {
	return &RateLimiter{
		rateLimit:  rateLimit,
		timeWindow: timeWindow,
		counters:   make(map[string]int),
		timestamps: make(map[string]time.Time),
	}
}

func (r *RateLimiter) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ip := req.RemoteAddr
		r.mu.Lock()
		counter, ok := r.counters[ip]
		if !ok {
			counter = 0
		}
		timestamp, ok := r.timestamps[ip]
		if !ok {
			timestamp = time.Now()
		}
		r.mu.Unlock()

		if time.Since(timestamp) >= r.timeWindow {
			r.mu.Lock()
			delete(r.counters, ip)
			delete(r.timestamps, ip)
			r.mu.Unlock()
		} else if counter >= r.rateLimit {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		} else {
			r.mu.Lock()
			r.counters[ip]++
			r.timestamps[ip] = time.Now()
			r.mu.Unlock()
		}
		next.ServeHTTP(w, req)
	})
}
