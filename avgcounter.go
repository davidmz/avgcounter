// Package avgcounter implements a simple EMA (Exponential Moving Average)
// counter. The New function creates a counter with the only parameter:
// avgInterval. Every Add call adds the value to the counter. The current value
// can be obtained using the Value method.
//
// The counter holds the exponentially (by time) weighted average of all added
// values.
//
package avgcounter

import (
	"math"
	"sync"
	"time"
)

var nowTime time.Time

func getNow() time.Time {
	if nowTime.IsZero() {
		return time.Now()
	}
	return nowTime
}

// Counter is an EMA (Exponential Moving Average) counter.
type Counter struct {
	lastValue   float64
	lastTime    time.Time
	avgInterval time.Duration
	lk          sync.RWMutex
}

// New creates a new Counter with the given avgInterval.
func New(avgInterval time.Duration) *Counter {
	return &Counter{avgInterval: avgInterval}
}

// Value returns the current value of the counter.
func (c *Counter) Value() float64 {
	c.lk.RLock()
	defer c.lk.RUnlock()
	return c.value()
}

// ValuePer returns the current value of the counter, normalized to the given
// interval. It is actually a Value() * interval / avgInterval.
func (c *Counter) ValuePer(interval time.Duration) float64 {
	return c.Value() * float64(interval) / float64(c.avgInterval)
}

// Add adds a new value to the counter.
func (c *Counter) Add(v float64) {
	c.lk.Lock()
	defer c.lk.Unlock()

	c.lastValue = c.value() + v
	c.lastTime = getNow()
}

// Private methods

func (c *Counter) value() float64 {
	if c.lastValue == 0 {
		return 0
	}
	return c.lastValue * math.Exp(-float64(getNow().Sub(c.lastTime))/float64(c.avgInterval))
}
