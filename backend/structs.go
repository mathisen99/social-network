package backend

import "time"

// Limiter struct
type Limiter struct {
	limit    int           // maximum number of events
	interval time.Duration // duration of the interval
	events   int           // number of events in the current interval
	reset    time.Time     // time when the interval resets
}
