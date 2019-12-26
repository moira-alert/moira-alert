package graphite

import "time"

// MetricsMap implements meter collection abstraction
type MetricsMap interface {
	AddMetric(name, path string)
	GetMetric(name string) (Meter, bool)
}

// Meter count events to produce exponentially-weighted moving average rates
// at one-, five-, and fifteen-minutes and a mean rate.
type Meter interface {
	Mark(int64)
}

// Timer capture the duration and rate of events.
type Timer interface {
	Count() int64
	UpdateSince(time.Time)
}

// Histogram calculate distribution statistics from a series of int64 values.
type Histogram interface {
	Update(int64)
}

// Counter hold an int64 value that can be incremented and decremented.
type Counter interface {
	Count() int64
	Inc(int64)
}
