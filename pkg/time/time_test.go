package time

import (
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
)

func TestFloat64Sec(t *testing.T) {
	assert.Equal(t, float64(60.07), Float64Sec(60070*time.Millisecond))
}

func TestMediaEpoch(t *testing.T) {
	// Calculate 1hr since MediaEpoch
	elapse, err := time.ParseDuration("3600s")
	assert.Nil(t, err)

	expected := time.Date(1904, time.January, 1, 1, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, MediaEpoch.Add(elapse))
}

func TestMediaTime(t *testing.T) {
	elapse := uint64(3457708564)

	mediaTime, err := MediaTime(elapse)
	assert.Nil(t, err)
	expected := time.Date(2013, time.July, 26, 18, 36, 4, 0, time.UTC)
	assert.Equal(t, expected, mediaTime)
}

func TestAge(t *testing.T) {
	{
		// 3 days
		then := time.Date(2018, time.May, 10, 1, 2, 3, 4, time.UTC)
		now := time.Date(2018, time.May, 13, 1, 2, 3, 4, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return now })
		assert.Equal(t, "3d", Age(then))
		patch.Unpatch()
	}
	{
		// 3 days and hours don't matter
		then := time.Date(2018, time.May, 10, 1, 2, 3, 4, time.UTC)
		now := time.Date(2018, time.May, 13, 10, 2, 3, 4, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return now })
		assert.Equal(t, "3d", Age(then))
		patch.Unpatch()
	}
	{
		// 9 hours
		then := time.Date(2018, time.May, 10, 1, 2, 3, 4, time.UTC)
		now := time.Date(2018, time.May, 10, 10, 2, 3, 4, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return now })
		assert.Equal(t, "9h", Age(then))
		patch.Unpatch()
	}
	{
		// 10 miniutes
		then := time.Date(2018, time.May, 10, 1, 2, 3, 4, time.UTC)
		now := time.Date(2018, time.May, 10, 1, 12, 3, 4, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return now })
		assert.Equal(t, "10m", Age(then))
		patch.Unpatch()
	}
	{
		// 10 seconds
		then := time.Date(2018, time.May, 10, 1, 2, 3, 4, time.UTC)
		now := time.Date(2018, time.May, 10, 1, 2, 13, 4, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return now })
		assert.Equal(t, "10s", Age(then))
		patch.Unpatch()
	}
	{
		// 0s as nanoseconds are not taken into account
		then := time.Date(2018, time.May, 10, 1, 2, 3, 4, time.UTC)
		now := time.Date(2018, time.May, 10, 1, 2, 3, 14, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return now })
		assert.Equal(t, "0s", Age(then))
		patch.Unpatch()
	}
}
