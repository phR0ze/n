// Package color provides basic terminal color output via ANSI Escape Codes
// https://misc.flogisoft.com/bash/tip_colors_and_formatting
package color

import (
	"fmt"
	"strings"

	"github.com/phR0ze/n/pkg/term"
)

const (
	gBold      = 1
	gUnderline = 4
	gEscape    = "\x1b"

	// None will reset colors
	None = 0
)

const (
	gBlack    = "30"
	gRed      = "31"
	gGreen    = "32"
	gYellow   = "33"
	gBlue     = "34"
	gMagenta  = "35"
	gCyan     = "36"
	gRedL     = "91"
	gGreenL   = "92"
	gYellowL  = "93"
	gBlueL    = "94"
	gMagentaL = "95"
	gCyanL    = "96"
	gWhite    = "97"
	gFuscia   = "38;5;201"
	gFusciaL  = "38;5;213"
	gGray     = "38;5;247" // "90"
	gGrayL    = "37"       // "37"
	gOrange   = "38;5;202"
	gOrangeL  = "38;5;208"
)

var (
	gIsTTy bool
)

func init() {
	if term.IsTTY() {
		gIsTTy = true
	}
}

// set the format attributes
func set(color string, bold int, underline int) string {
	if !gIsTTy {
		return ""
	}

	f := []string{}
	a := []interface{}{gEscape}

	// Set Bold
	if bold == 1 {
		f = append(f, "%v")
		a = append(a, gBold)
	}

	// Set underline
	if underline == 1 {
		f = append(f, "%v")
		a = append(a, gUnderline)
	}

	// Set color
	if color != "" {
		f = append(f, "%v")
		a = append(a, color)
	}

	// Format
	if len(f) > 0 {
		seq := strings.Join([]string{"%s[", strings.Join(f, ";"), "m"}, "")
		return fmt.Sprintf(seq, a...)
	}
	return ""
}

// reset all formatting attributes
func reset() string {
	if !gIsTTy {
		return ""
	}
	return fmt.Sprintf("%s[%dm", gEscape, None)
}

// Blue
// -------------------------------------------------------------------------------------------------

// Blue colors the string blue
func Blue(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlue, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// BlueB colors the string blue and bolds it
func BlueB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlue, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// BlueL colors the string light blue
func BlueL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlueL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// BlueU colors the string blue and underlines it
func BlueU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlueL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// BlueLU colors the string light blue and underlines it
func BlueLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlueL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// BlueBL colors the string light blue and bolds it
func BlueBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlueL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// BlueBU colors the string blue and bolds and underlines it
func BlueBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlue, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// BlueBLU colors the string light blue and bolds and underlines it
func BlueBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gBlueL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Cyan
// -------------------------------------------------------------------------------------------------

// Cyan colors the string cyan
func Cyan(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyan, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// CyanB colors the string cyan and bolds it
func CyanB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyan, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// CyanL colors the string light cyan
func CyanL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyanL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// CyanU colors the string cyan and underlines it
func CyanU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyanL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// CyanLU colors the string light cyan and underlines it
func CyanLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyanL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// CyanBL colors the string light cyan and bolds it
func CyanBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyanL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// CyanBU colors the string cyan and bolds and underlines it
func CyanBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyan, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// CyanBLU colors the string light cyan and bolds and underlines it
func CyanBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gCyanL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Fuscia
// -------------------------------------------------------------------------------------------------

// Fuscia colors the string fuscia
func Fuscia(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFuscia, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// FusciaB colors the string fuscia and bolds it
func FusciaB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFuscia, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// FusciaL colors the string light fuscia
func FusciaL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFusciaL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// FusciaU colors the string fuscia and underlines it
func FusciaU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFusciaL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// FusciaLU colors the string light fuscia and underlines it
func FusciaLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFusciaL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// FusciaBL colors the string light fuscia and bolds it
func FusciaBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFusciaL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// FusciaBU colors the string fuscia and bolds and underlines it
func FusciaBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFuscia, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// FusciaBLU colors the string light fuscia and bolds and underlines it
func FusciaBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gFusciaL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Gray
// -------------------------------------------------------------------------------------------------

// Gray colors the string gray
func Gray(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGray, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GrayB colors the string gray and bolds it
func GrayB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGray, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GrayL colors the string light gray
func GrayL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGrayL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GrayU colors the string gray and underlines it
func GrayU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGrayL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// GrayLU colors the string light gray and underlines it
func GrayLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGrayL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// GrayBL colors the string light gray and bolds it
func GrayBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGrayL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GrayBU colors the string gray and bolds and underlines it
func GrayBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGray, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// GrayBLU colors the string light gray and bolds and underlines it
func GrayBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGrayL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Green
// -------------------------------------------------------------------------------------------------

// Green colors the string green
func Green(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreen, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GreenB colors the string green and bolds it
func GreenB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreen, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GreenL colors the string light green
func GreenL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreenL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GreenU colors the string green and underlines it
func GreenU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreenL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// GreenLU colors the string light green and underlines it
func GreenLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreenL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// GreenBL colors the string light green and bolds it
func GreenBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreenL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// GreenBU colors the string green and bolds and underlines it
func GreenBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreen, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// GreenBLU colors the string light green and bolds and underlines it
func GreenBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gGreenL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Magenta
// -------------------------------------------------------------------------------------------------

// Magenta colors the string magenta
func Magenta(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagenta, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// MagentaB colors the string magenta and bolds it
func MagentaB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagenta, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// MagentaL colors the string light magenta
func MagentaL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagentaL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// MagentaU colors the string magenta and underlines it
func MagentaU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagentaL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// MagentaLU colors the string light magenta and underlines it
func MagentaLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagentaL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// MagentaBL colors the string light magenta and bolds it
func MagentaBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagentaL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// MagentaBU colors the string magenta and bolds and underlines it
func MagentaBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagenta, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// MagentaBLU colors the string light magenta and bolds and underlines it
func MagentaBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gMagentaL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Orange
// -------------------------------------------------------------------------------------------------

// Orange colors the string orange
func Orange(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrange, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// OrangeB colors the string orange and bolds it
func OrangeB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrange, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// OrangeL colors the string light orange
func OrangeL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrangeL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// OrangeU colors the string orange and underlines it
func OrangeU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrangeL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// OrangeLU colors the string light orange and underlines it
func OrangeLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrangeL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// OrangeBL colors the string light orange and bolds it
func OrangeBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrangeL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// OrangeBU colors the string orange and bolds and underlines it
func OrangeBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrange, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// OrangeBLU colors the string light orange and bolds and underlines it
func OrangeBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gOrangeL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Red
// -------------------------------------------------------------------------------------------------

// Red colors the string red
func Red(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRed, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// RedB colors the string red and bolds it
func RedB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRed, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// RedL colors the string light red
func RedL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRedL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// RedU colors the string red and underlines it
func RedU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRedL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// RedLU colors the string light red and underlines it
func RedLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRedL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// RedBL colors the string light red and bolds it
func RedBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRedL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// RedBU colors the string red and bolds and underlines it
func RedBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRed, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// RedBLU colors the string light red and bolds and underlines it
func RedBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gRedL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// Yellow
// -------------------------------------------------------------------------------------------------

// Yellow colors the string yellow
func Yellow(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellow, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// YellowB colors the string yellow and bolds it
func YellowB(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellow, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// YellowL colors the string light yellow
func YellowL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellowL, 0, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// YellowU colors the string yellow and underlines it
func YellowU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellowL, 0, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// YellowLU colors the string light yellow and underlines it
func YellowLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellowL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// YellowBL colors the string light yellow and bolds it
func YellowBL(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellowL, 1, 0), fmt.Sprintf(format, a...), reset()}, "")
}

// YellowBU colors the string yellow and bolds and underlines it
func YellowBU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellow, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}

// YellowBLU colors the string light yellow and bolds and underlines it
func YellowBLU(format string, a ...interface{}) string {
	return strings.Join([]string{set(gYellowL, 1, 1), fmt.Sprintf(format, a...), reset()}, "")
}
