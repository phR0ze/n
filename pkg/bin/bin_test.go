package bin

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMediaDuration(t *testing.T) {
	scaleTime := uint32(600)

	// 3600 sec duration in BigEndian format
	data := []byte{0x00, 0x00, 0x8c, 0xca}

	duration := MediaDuration32BE(data, scaleTime)
	assert.Equal(t, 1*time.Minute, duration)
}

func TestPutUint32BE(t *testing.T) {
	name := make([]byte, 4)
	PutUint32BE(name, uint32(0x74726578))

	assert.Equal(t, []byte{0x74, 0x72, 0x65, 0x78}, name)
}

func TestFixed32BE(t *testing.T) {

	// Data in fixedpoint (i.e. float) BigEndian format
	data := []byte{0x05, 0x00, 0x00, 0x00}

	// Incorrectly convert it straight to a uint32
	assert.Equal(t, uint32(83886080), Uint32BE(data))

	// Correctly convert it to float
	assert.Equal(t, float64(1280), Fixed32BE(data))
}

func TestUint16BE(t *testing.T) {

	// Data in BigEndian format
	data := []byte{0x00, 0x05}

	assert.Equal(t, uint16(5), Uint16BE(data))
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
