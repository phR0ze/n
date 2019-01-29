package bin

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFixed32BE(t *testing.T) {

	// Data in fixedpoint (i.e. float) BigEndian format
	data := []byte{0x05, 0x00, 0x00, 0x00}

	// Incorrectly convert it straight to a uint32
	assert.Equal(t, uint32(83886080), Uint32BE(data))

	// Correctly convert it to float
	assert.Equal(t, float64(1280), Fixed32BE(data))
}

func TestUint32BE(t *testing.T) {

	// Data in BigEndian format
	data := []byte{0x05, 0x00, 0x00, 0x00}

	assert.Equal(t, uint32(83886080), Uint32BE(data))
}

func TestMediaTime32BE(t *testing.T) {
	// Data in media time BigEndian format
	data := []byte{0xce, 0x18, 0x72, 0x18}

	expected := time.Date(2013, time.July, 26, 18, 36, 8, 0, time.UTC)
	assert.Equal(t, expected, MediaTime32BE(data))
}
