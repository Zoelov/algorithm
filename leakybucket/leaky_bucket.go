package leakybucket

import (
	"sync/atomic"
	"time"
)

type leakyBucket struct {
	capacity int64
	rate     float64
	lastTime atomic.Uint64
	current  atomic.Uint64
}

func New(capacity int64, rate float64) *leakyBucket {
	return &leakyBucket{
		lastTime: atomic.Uint64{},
		current:  atomic.Uint64{},
		capacity: capacity,
		rate:     rate / 1000,
	}
}

func (l *leakyBucket) Allow() bool {
	now := time.Now().UnixNano() / 1e6

	for {
		lastTime := l.lastTime.Load()
		oldWater := l.current.Load()
		var leakedWater int64
		if now > int64(lastTime) {
			leakedWater = int64(float64(now-int64(lastTime)) * l.rate)
		}

		newWater := max(int64(oldWater)-leakedWater, 0)

		if newWater >= l.capacity*1000 {
			return false
		}

		if l.current.CompareAndSwap(oldWater, uint64(newWater+1000)) {
			l.lastTime.CompareAndSwap(lastTime, uint64(now))
			return true
		}

	}

}
