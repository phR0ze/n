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
func Age(other time.Time) (result string) {
	duration := time.Since(other)

	// Determin how many days elapsed
	days := int(duration.Hours()) / 24
	if days > 0 {
		result = fmt.Sprintf("%dd", days)
		duration = duration - time.Duration(days*24)*time.Hour
	}

	// Determin how many hours elapsed beyond days
	hours := int(duration.Hours())
	if hours > 0 {
		result = result + fmt.Sprintf("%dh", hours)
		duration = duration - time.Duration(hours)*time.Hour
	}

	// Determin how many mins elapsed beyond hours
	mins := int(duration.Minutes())
	if mins > 0 {
		result = result + fmt.Sprintf("%dm", mins)
		duration = duration - time.Duration(mins)*time.Minute
	}

	// Determin how many secs elapsed beyond mins
	secs := int(duration.Seconds())
	if secs > 0 {
		result = result + fmt.Sprintf("%ds", secs)
		duration = duration - time.Duration(secs)*time.Second
	}

	// Default to 0s
	if result == "" {
		result = "0s"
	}
	return
}
