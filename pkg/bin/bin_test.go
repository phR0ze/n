package bin

import (
	"testing"
	"time"

	ntime "github.com/phR0ze/n/pkg/time"
	"github.com/stretchr/testify/assert"
)

func TestMediaDuration(t *testing.T) {
	timeScale := uint32(600)

	// 3600 sec duration in BigEndian format
	data := []byte{0x00, 0x00, 0x8c, 0xca}
	assert.Equal(t, uint32(36042), Uint32BE(data))

	// Ensure float accuracy of time scale divsor
	assert.Equal(t, uint32(60), uint32(36042)/timeScale)
	assert.Equal(t, float64(60.07), float64(36042)/float64(timeScale))

	// Now keep the float precision with a duration
	duration := time.Millisecond * time.Duration(float64(Uint32BE(data))/float64(timeScale)*1000.0)
	assert.Equal(t, 60070*time.Millisecond, duration)
	assert.Equal(t, float64(60.07), float64(duration)/float64(time.Millisecond)/1000.0)

	// Now test with media duration
	duration = MediaDuration32BE(data, timeScale)
	assert.Equal(t, float64(60.07), ntime.Float64Sec(duration))
	assert.Equal(t, 1*time.Minute+70*time.Millisecond, duration)
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

func TestInt16BE(t *testing.T) {

	// Data in BigEndian format
	data := []byte{0x00, 0x05}

	assert.Equal(t, 5, Int16BE(data))
}

func TestUint16BE(t *testing.T) {

	// Data in BigEndian format
	data := []byte{0x00, 0x05}

	assert.Equal(t, uint16(5), Uint16BE(data))
}

func TestInt32BE(t *testing.T) {

	// Data in BigEndian format
	data := []byte{0x05, 0x00, 0x00, 0x00}

	assert.Equal(t, 83886080, Int32BE(data))
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
