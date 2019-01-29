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

// Uint32BE reads data as BigEndian
func Uint32BE(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

// MediaTime32BE reads the data as BigEndian and convert to time
func MediaTime32BE(data []byte) time.Time {
	sec := uint64(binary.BigEndian.Uint32(data))
	result, _ := ntime.MediaTime(sec)
	return result
}

// Fixed32BE reads data as BigEndian fixed point 16.16 and converts to float
func Fixed32BE(data []byte) float64 {

	// 2 bytes or 16 bits are used for the whole number part of the value
	whole := float64(binary.BigEndian.Uint16(data[:2]))

	// 2 bytes or 16 bits are used for the fraction part of the value
	fraction := float64(binary.BigEndian.Uint16(data[2:]))

	return whole + fraction/Fixed16bitDiv
}
