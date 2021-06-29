package pool

import (
	"sync"
	"time"
)

var timerPool sync.Pool
// TODO copy from gorpc
// AcquireTimer
func AcquireTimer(timeout time.Duration) *time.Timer {
	tv := timerPool.Get()
	if tv == nil {
		return time.NewTimer(timeout)
	}

	t := tv.(*time.Timer)
	if t.Reset(timeout) {
		panic("")
	}
	return t
}

// ReleaseTimer
func ReleaseTimer(t *time.Timer) {
	if !t.Stop() {
		// Collect possibly added time from the channel
		// if timer has been stopped and nobody collected its' value.
		select {
		case <-t.C:
		default:
		}
	}

	timerPool.Put(t)
}
