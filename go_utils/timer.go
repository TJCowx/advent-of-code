package go_utils

import (
	"fmt"
	"time"
)

// Simple utility to measure the start and end of a process
type Timer struct {
	start time.Time
	end   time.Time
}

// Sets the start of the timer to the current time
// Run this before end
func (t *Timer) Start() {
	t.start = time.Now()
}

// Sets the end of the timer to the current time
// Run this after start
func (t *Timer) End() {
	t.end = time.Now()
}

// Returns a formatted string displaying how long the time has lapsed
// Example output: "Timer ran for: 1.234567s"
func (t *Timer) TimeLapsed() string {
	return fmt.Sprintf("Timer ran for: %s", t.end.Sub(t.start))
}
