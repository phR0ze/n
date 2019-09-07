// Package opt provides a simple struct for options that can be passed in as an optional last
// variadic parameter.
//
// Opt was created with the intent of promoting a pattern for handling optional parameters to
// functions. By making a parameter optional we effectively proved a third value of unset for
// a given param which other patterns don't have. For instance by passing in a struct as
// options, you get the value and Go's default if not set, but you don't know if the not set
// value was intentional where as with an option not existing know the user intended to not set
// that option. This is an important distinction that the struct based option pattern does not
// have.
//
// Options Pattern:
//
// Packages making use of the variadic options pattern should ignore options that are not
// supported and clearly call out in function comments which options are supported.
//
// Create helper functions that wrap your constants to allow the compiler to assist. The creation
// functions should be called <option>Opt, getters Get<option>Opt and default helper
// Default<option>Opt. e.g. DebugOpt, GetDebugOpt and DefaultDebugOpt
//
// Properties Pattern:
//
// Custom types often need similar properties and the <properties>Props helper structs
package opt

import (
	"io"
	"os"
)

// StdProps provides common properties for custom types
type StdProps struct {
	StdStreamProps
	Home    string // Home path to use
	Quiet   bool   // Quiet mode when true
	Debug   bool   // Debug mode when true
	DryRun  bool   // Dryrun mode when true
	Testing bool   // Testing mode when true
}

// StdStreamProps provides common io stream properties for custom types
type StdStreamProps struct {
	In  io.Reader // Input stream to use
	Out io.Writer // Output stream to use
	Err io.Writer // Error stream to use
}

// Opt provides a mechanism for passing in optional paramaters to functions.
type Opt struct {
	Key string      // name of the option
	Val interface{} // value of the option
}

// Add an option to the options slice if it doesn't exist.
// Just an alias to Default, but the alternate naming is more intuitive in
// different scenarios.
func Add(opts *[]*Opt, opt *Opt) bool {
	if opts == nil || opt == nil || Exists(*opts, opt.Key) {
		return false
	}
	*opts = append(*opts, opt)
	return true
}

// Copy the given options slice removing nil options. Although the slice and
// options are new distinct objects the Values are the original objects.
func Copy(opts []*Opt) []*Opt {
	newOpts := []*Opt{}
	for _, o := range opts {
		if o != nil {
			newOpts = append(newOpts, &Opt{Key: o.Key, Val: o.Val})
		}
	}
	return newOpts
}

// Default adds an option to the options slice if it doesn't exist.
// Returns true the option was added to the options slice or false if
// the given slice or option are nil or the option already exists in the slice.
func Default(opts *[]*Opt, opt *Opt) bool {
	if opts == nil || opt == nil || Exists(*opts, opt.Key) {
		return false
	}
	*opts = append(*opts, opt)
	return true
}

// Exists checks if the given opt exists in the opts slice by key
func Exists(opts []*Opt, key string) bool {
	return Get(opts, key) != nil
}

// Get an option by the given key
func Get(opts []*Opt, key string) *Opt {
	for _, o := range opts {
		if o != nil && o.Key == key {
			return o
		}
	}
	return nil
}

// Overwrite replaces an existing option or adds the option if it doesn't exist.
func Overwrite(opts *[]*Opt, opt *Opt) {
	if opts != nil && opt != nil {
		if Exists(*opts, opt.Key) {
			Get(*opts, opt.Key).Val = opt.Val
		} else {
			Add(opts, opt)
		}
	}
}

// Remove an existing option from the options slice if found
func Remove(opts *[]*Opt, opt *Opt) {
	if opts != nil && opt != nil {
		for i := len(*opts) - 1; i >= 0; i-- {
			if (*opts)[i] != nil && (*opts)[i].Key == opt.Key {
				if i+1 < len(*opts) {
					*opts = append((*opts)[:i], (*opts)[i+1:]...)
				} else {
					*opts = (*opts)[:i]
				}
				return
			}
		}
	}
}

// In Option
// -------------------------------------------------------------------------------------------------

// InOpt creates the new option with the given value
func InOpt(val io.Reader) *Opt {
	return &Opt{Key: "in", Val: val}
}

// GetInOpt finds and returns the option's value or defaults to os.Stdin
func GetInOpt(opts []*Opt) io.Reader {
	if o := Get(opts, "in"); o != nil {
		if val, ok := o.Val.(io.Reader); ok {
			return val
		}
	}
	return os.Stdin
}

// DefaultInOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultInOpt(opts *[]*Opt, val io.Reader) io.Reader {
	if opts == nil || !Exists(*opts, "in") {
		return val
	}
	return GetInOpt(*opts)
}

// Out Option
// -------------------------------------------------------------------------------------------------

// OutOpt creates the new option with the given value
func OutOpt(val io.Writer) *Opt {
	return &Opt{Key: "out", Val: val}
}

// GetOutOpt finds and returns the option's value or defaults to os.Stdout
func GetOutOpt(opts []*Opt) io.Writer {
	if o := Get(opts, "out"); o != nil {
		if val, ok := o.Val.(io.Writer); ok {
			return val
		}
	}
	return os.Stdout
}

// DefaultOutOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultOutOpt(opts *[]*Opt, val io.Writer) io.Writer {
	if opts == nil || !Exists(*opts, "out") {
		return val
	}
	return GetOutOpt(*opts)
}

// Err Option
// -------------------------------------------------------------------------------------------------

// ErrOpt creates the new option with the given value
func ErrOpt(val io.Writer) *Opt {
	return &Opt{Key: "err", Val: val}
}

// GetErrOpt finds and returns the option's value or defaults to os.Stderr
func GetErrOpt(opts []*Opt) io.Writer {
	if o := Get(opts, "err"); o != nil {
		if val, ok := o.Val.(io.Writer); ok {
			return val
		}
	}
	return os.Stderr
}

// DefaultErrOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultErrOpt(opts *[]*Opt, val io.Writer) io.Writer {
	if opts == nil || !Exists(*opts, "err") {
		return val
	}
	return GetErrOpt(*opts)
}

// Home Option
// -------------------------------------------------------------------------------------------------

// HomeOpt creates the new option with the given value
func HomeOpt(val string) *Opt {
	return &Opt{Key: "home", Val: val}
}

// GetHomeOpt finds and returns the option's value or defaults to empty string
func GetHomeOpt(opts []*Opt) string {
	if o := Get(opts, "home"); o != nil {
		if val, ok := o.Val.(string); ok {
			return val
		}
	}
	return ""
}

// DefaultHomeOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultHomeOpt(opts *[]*Opt, val string) string {
	if opts == nil || !Exists(*opts, "home") {
		return val
	}
	return GetHomeOpt(*opts)
}

// Quiet Option
//--------------------------------------------------------------------------------------------------

// QuietOpt creates the new option with the given value
func QuietOpt(val bool) *Opt {
	return &Opt{Key: "quiet", Val: val}
}

// GetQuietOpt finds and returns the option's value or defaults to false
func GetQuietOpt(opts []*Opt) (result bool) {
	if option := Get(opts, "quiet"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultQuietOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultQuietOpt(opts *[]*Opt, val bool) bool {
	if opts == nil || !Exists(*opts, "quiet") {
		return val
	}
	return GetQuietOpt(*opts)
}

// Debug Option
//--------------------------------------------------------------------------------------------------

// DebugOpt creates the new option with the given value
func DebugOpt(val bool) *Opt {
	return &Opt{Key: "debug", Val: val}
}

// GetDebugOpt finds and returns the option's value or defaults to false
func GetDebugOpt(opts []*Opt) (result bool) {
	if option := Get(opts, "debug"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultDebugOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultDebugOpt(opts *[]*Opt, val bool) bool {
	if opts == nil || !Exists(*opts, "debug") {
		return val
	}
	return GetDebugOpt(*opts)
}

// DryRun Option
//--------------------------------------------------------------------------------------------------

// DryRunOpt creates the new option with the given value
func DryRunOpt(val bool) *Opt {
	return &Opt{Key: "dry-run", Val: val}
}

// GetDryRunOpt finds and returns the option's value or defaults to false
func GetDryRunOpt(opts []*Opt) (result bool) {
	if option := Get(opts, "dry-run"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultDryRunOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultDryRunOpt(opts *[]*Opt, val bool) bool {
	if opts == nil || !Exists(*opts, "dry-run") {
		return val
	}
	return GetDryRunOpt(*opts)
}

// Testing Option
//--------------------------------------------------------------------------------------------------

// TestingOpt creates the new option with the given value
func TestingOpt(val bool) *Opt {
	return &Opt{Key: "testing", Val: val}
}

// GetTestingOpt finds and returns the option's value or defaults to false
func GetTestingOpt(opts []*Opt) (result bool) {
	if option := Get(opts, "testing"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultTestingOpt returns the option value if found else the one given
// Use this when the Get's default is not desirable.
func DefaultTestingOpt(opts *[]*Opt, val bool) bool {
	if opts == nil || !Exists(*opts, "testing") {
		return val
	}
	return GetTestingOpt(*opts)
}
