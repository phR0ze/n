// Package unit provides some helpful unit conversions
package unit

import (
	"fmt"
	"strings"
)

const (
	Kibibyte = 1024
	KiB      = Kibibyte
	Mebibyte = Kibibyte * 1024
	MiB      = Mebibyte
	Gibibyte = Mebibyte * 1024
	GiB      = Gibibyte
	Tebibyte = Gibibyte * 1024
	TiB      = Tebibyte
)

// ToBase2Bytes converts the given value in bytes to a human readable format
func ToBase2Bytes(val int64) (result string) {
	unit := "bytes"
	value := float64(val)

	switch {
	case val >= TiB:
		value = value / TiB
		unit = "TiB"
	case val >= GiB:
		value = value / GiB
		unit = "GiB"
	case val >= MiB:
		value = value / MiB
		unit = "MiB"
	case val >= KiB:
		value = value / KiB
		unit = "KiB"
	}

	result = strings.TrimSuffix(fmt.Sprintf("%.2f", value), ".00")
	result = fmt.Sprintf("%s %s", result, unit)
	return
}
