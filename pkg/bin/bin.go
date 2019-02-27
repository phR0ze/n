// Package bin provides some low level binary protocol helpers
package bin

import (
	"encoding/binary"
	"time"

	ntime "github.com/phR0ze/n/pkg/time"
)

const (
	// Fixed16bitDiv is the 2^16 value used to convert fractional part to fraction
	Fixed16bitDiv = 65536.0
)

// PutUint32BE is just a wrapper around binary.BigEndian for convenience
func PutUint32BE(data []byte, val uint32) {
	binary.BigEndian.PutUint32(data, val)
}

// Int8BE reads data as BigEndian
func Int8BE(data byte) int {
	return int(uint8(data))
}

// Int16BE reads data as BigEndian
func Int16BE(data []byte) int {
	return int(binary.BigEndian.Uint16(data))
}

// Uint16BE reads data as BigEndian
func Uint16BE(data []byte) uint16 {
	return binary.BigEndian.Uint16(data)
}

// Int32BE reads data as BigEndian
func Int32BE(data []byte) int {
	return int(binary.BigEndian.Uint32(data))
}

// Uint32BE reads data as BigEndian
func Uint32BE(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

// Uint64BE reads data as BigEndian
func Uint64BE(data []byte) uint64 {
	return binary.BigEndian.Uint64(data)
}

// MediaDuration32BE reads 4 bytes of data as BigEndian and converts it to
// a duration taking into account the time scale. data is in media units
func MediaDuration32BE(data []byte, timeScale uint32) time.Duration {
	nanosec := float64(Uint32BE(data)) / float64(timeScale) * 1000000000.0
	return time.Duration(nanosec)
}

// MediaTime32BE reads 4 bytes of data as BigEndian and convert to time
func MediaTime32BE(data []byte) time.Time {
	sec := uint64(binary.BigEndian.Uint32(data))
	result, _ := ntime.MediaTime(sec)
	return result
}

// Fixed32BE reads 4 bytes of data as BigEndian fixed point 16.16 and converts to float
func Fixed32BE(data []byte) float64 {

	// 2 bytes or 16 bits are used for the whole number part of the value
	whole := float64(binary.BigEndian.Uint16(data[:2]))

	// 2 bytes or 16 bits are used for the fraction part of the value
	fraction := float64(binary.BigEndian.Uint16(data[2:]))

	return whole + fraction/Fixed16bitDiv
}
