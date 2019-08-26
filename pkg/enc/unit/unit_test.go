package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBase2Bytes(t *testing.T) {
	assert.Equal(t, "10 bytes", HumanBase2(int64(10)))
	assert.Equal(t, "1 KiB", HumanBase2(int64(1024)))
	assert.Equal(t, "4.91 KiB", HumanBase2(int64(5024)))
	assert.Equal(t, "3 TiB", HumanBase2(int64(3*TiB)))
	assert.Equal(t, "3 GiB", HumanBase2(int64(3*GiB+500)))
	assert.Equal(t, "3.49 GiB", HumanBase2(int64(3*GiB+500*MiB)))
	assert.Equal(t, "3.05 MiB", HumanBase2(int64(3*MiB+50000)))
	assert.Equal(t, "3.05 MiB", HumanBase2(int64(3195728)))
}

func TestToKiB(t *testing.T) {
	assert.Equal(t, 1000.0, ToKiB(1000*KiB))
}

func TestToMiB(t *testing.T) {
	assert.Equal(t, 3.0476837158203125, ToMiB(3195728))
}

func TestToGiB(t *testing.T) {
	assert.Equal(t, 1000.0, ToGiB(1000*GiB))
}

func TestToTiB(t *testing.T) {
	assert.Equal(t, 1000.0, ToTiB(1000*TiB))
}
