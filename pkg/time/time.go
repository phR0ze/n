// Package time provides simple time helper functions
package time

import (
	"fmt"
	"time"
)

var (
	// MediaEpoch provides a time object for midnight, January 1, 1904
	MediaEpoch = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)
)

// Float64Sec converts a duration into a float64 seconds without loosing precision
func Float64Sec(duration time.Duration) float64 {
	return float64(duration) / float64(time.Millisecond) / 1000.0
}

// FromIntSec converts the given integer in seconds to a time.Duration
func FromIntSec(sec int) time.Duration {
	return time.Duration(int64(sec) * int64(time.Second))
}

// MediaTime calculates the time given the elapse in seconds since MediaEpoch
func MediaTime(sec uint64) (result time.Time, err error) {
	dura := time.Second * time.Duration(sec)
	result = MediaEpoch.Add(dura)
	return
}

// Age calculates the elapse time in days from a time.Time object
func Age(other time.Time) string {
	days := int(time.Since(other).Hours()) / 24
	hours := int(time.Since(other).Hours())
	mins := int(time.Since(other).Minutes())
	secs := int(time.Since(other).Seconds())

	switch {
	case days > 0:
		return fmt.Sprintf("%dd", days)
	case hours > 0:
		return fmt.Sprintf("%dh", hours)
	case mins > 0:
		return fmt.Sprintf("%dm", mins)
	case secs > 0:
		return fmt.Sprintf("%ds", secs)
	}
	return "0s"
}
