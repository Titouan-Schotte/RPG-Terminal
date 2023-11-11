package utils

import "time"

func Timeout(seconds float64) {
	duration := time.Duration(seconds * float64(time.Second))
	time.Sleep(duration)
}
