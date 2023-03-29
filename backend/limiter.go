package backend

import (
	"fmt"
	"net/http"
	misc "social-network/misc"
	"time"
)

// The Rate limiter now it is set to 50 requests per second
var limiter = NewLimiter(50, 1*time.Second)

// Limit is a middleware that limits the number of requests per second
func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			fmt.Fprint(w, "429 Too Many Requests")
			return
		}
		// If the limit has not been reached, then continue to server our regular handler
		next.ServeHTTP(w, r)
	})
}

// Sets the limit and interval
func NewLimiter(limit int, interval time.Duration) *Limiter {
	return &Limiter{
		limit:    limit,
		interval: interval,
		reset:    time.Now().Add(interval),
	}
}

// Checks if the limit has been reached
func (l *Limiter) Allow() bool {
	now := time.Now()

	if now.After(l.reset) {
		l.events = 1
		l.reset = now.Add(l.interval)
		return true
	}

	l.events++
	if l.events > l.limit {
		fmt.Println(misc.Red, "Server -> Rate limit exceeded", misc.Reset)
		return false
	}

	return true
}
