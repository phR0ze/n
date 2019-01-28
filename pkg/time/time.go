package time

import (
	"fmt"
	"time"
)

var (
	// MediaEpoch provides a time object for midnight, January 1, 1904
	MediaEpoch = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)
)

// MediaTime calculates the time given the elapse in seconds since MediaEpoch
func MediaTime(sec uint32) (result time.Time, err error) {
	var dura time.Duration
	if dura, err = time.ParseDuration(fmt.Sprintf("%ds", sec)); err != nil {
		return
	}
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
