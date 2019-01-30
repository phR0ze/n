package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBase2Bytes(t *testing.T) {
	assert.Equal(t, "10 bytes", ToBase2Bytes(int64(10)))
	assert.Equal(t, "1 KiB", ToBase2Bytes(int64(1024)))
	assert.Equal(t, "4.91 KiB", ToBase2Bytes(int64(5024)))
	assert.Equal(t, "3 TiB", ToBase2Bytes(int64(3*TiB)))
	assert.Equal(t, "3 GiB", ToBase2Bytes(int64(3*GiB+500)))
	assert.Equal(t, "3.49 GiB", ToBase2Bytes(int64(3*GiB+500*MiB)))
	assert.Equal(t, "3.05 MiB", ToBase2Bytes(int64(3*MiB+50000)))
}
