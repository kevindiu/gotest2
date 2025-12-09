package main

import "time"

// SlowFunction sleeps for a duration.
func SlowFunction(d time.Duration) {
	time.Sleep(d)
}
