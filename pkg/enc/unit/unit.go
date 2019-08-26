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

// HumanBase2 converts the given value in bytes to a human readable format
// e.g. 3195728 = 3.05 MiB
func HumanBase2(val int64) (result string) {
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

// ToKiB converts the given value in bytes to increments of KiB
func ToKiB(val int64) (kib float64) {
	return float64(val) / float64(KiB)
}

// ToMiB converts the given value in bytes to increments of MiB
func ToMiB(val int64) (mib float64) {
	return float64(val) / float64(MiB)
}

// ToGiB converts the given value in bytes to increments of GiB
func ToGiB(val int64) (gib float64) {
	return float64(val) / float64(GiB)
}

// ToTiB converts the given value in bytes to increments of TiB
func ToTiB(val int64) (tib float64) {
	return float64(val) / float64(TiB)
}
