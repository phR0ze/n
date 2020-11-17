// Package opt provides common properties for custom types and support for the options pattern.
//
// Opt was created with the intent of promoting a pattern for handling optional parameters to
// functions. By making a parameter optional we effectively provide a third value of unset for
// a given param which other patterns don't have. For instance by passing in a struct as
// options, you get the value and Go's default if not set, but you don't know if the not set
// value was intentional where as with an option not existing you know the user intended to not
// set that option. This is an important distinction that the struct based option pattern does
// not have.
//
// Options Pattern:
//
// Packages making use of the variadic options pattern should ignore options that are not
// supported and clearly call out in function comments which options are supported.
//
// Create helper functions that wrap your constants to allow the compiler to assist. The creation
// functions should be called <option>Opt, checkers <option>OptExists, getters Get<option>Opt
// and default helper Default<option>Opt. e.g. DebugOpt, DebugOptExists, GetDebugOpt and
// DefaultDebugOpt
//
// Common Properties:
//
// To support the options pattern I've included some common reusable properties that custom types
// can inherit from and their supporting option helper functions.
package opt

import (
	"io"
	"os"
)

// Common options keys
var (
	InOptKey      = "in"
	OutOptKey     = "out"
	ErrOptKey     = "err"
	HomeOptKey    = "home"
	QuietOptKey   = "quiet"
	DebugOptKey   = "debug"
	DryRunOptKey  = "dry-run"
	TestingOptKey = "testing"
)

// Std provides an interface for the standard properties
type Std interface {
	StdStream
	Home(home ...string) string   // Home path to use
	Quiet(quiet ...bool) bool     // Quiet mode when true
	Debug(debug ...bool) bool     // Debug mode when true
	DryRun(dryrun ...bool) bool   // Dryrun mode when true
	Testing(testing ...bool) bool // Testing mode when true
}

// StdStream provides an interface for the standard streams
type StdStream interface {
	In(in ...io.Reader) io.Reader   // Input stream to use
	Out(out ...io.Writer) io.Writer // Output stream to use
	Err(err ...io.Writer) io.Writer // Error stream to use
}

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
func Overwrite(opts *[]*Opt, opt *Opt) *Opt {
	if opts != nil && opt != nil {
		if Exists(*opts, opt.Key) {
			Get(*opts, opt.Key).Val = opt.Val
		} else {
			Add(opts, opt)
		}
	}
	return opt
}

// Remove an existing option by key from the options list
func Remove(opts *[]*Opt, key string) {
	if opts != nil && key != "" {
		for i := len(*opts) - 1; i >= 0; i-- {
			if (*opts)[i] != nil && (*opts)[i].Key == key {
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
	return &Opt{Key: InOptKey, Val: val}
}

// InOptExists determines if the option exists in the given options
func InOptExists(opts []*Opt) bool {
	return Exists(opts, InOptKey)
}

// GetInOpt finds and returns the option's value or defaults to os.Stdin
func GetInOpt(opts []*Opt) io.Reader {
	if o := Get(opts, InOptKey); o != nil {
		if val, ok := o.Val.(io.Reader); ok {
			return val
		}
	}
	return os.Stdin
}

// DefaultInOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultInOpt(opts []*Opt, val io.Reader) io.Reader {
	if !Exists(opts, InOptKey) {
		return val
	}
	return GetInOpt(opts)
}

// OverwriteInOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteInOpt(opts *[]*Opt, val io.Reader) io.Reader {
	return Overwrite(opts, InOpt(val)).Val.(io.Reader)
}

// Out Option
// -------------------------------------------------------------------------------------------------

// OutOpt creates the new option with the given value
func OutOpt(val io.Writer) *Opt {
	return &Opt{Key: OutOptKey, Val: val}
}

// OutOptExists determines if the option exists in the given options
func OutOptExists(opts []*Opt) bool {
	return Exists(opts, OutOptKey)
}

// GetOutOpt finds and returns the option's value or defaults to os.Stdout
func GetOutOpt(opts []*Opt) io.Writer {
	if o := Get(opts, OutOptKey); o != nil {
		if val, ok := o.Val.(io.Writer); ok {
			return val
		}
	}
	return os.Stdout
}

// DefaultOutOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultOutOpt(opts []*Opt, val io.Writer) io.Writer {
	if !Exists(opts, OutOptKey) {
		return val
	}
	return GetOutOpt(opts)
}

// OverwriteOutOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteOutOpt(opts *[]*Opt, val io.Writer) io.Writer {
	return Overwrite(opts, OutOpt(val)).Val.(io.Writer)
}

// Err Option
// -------------------------------------------------------------------------------------------------

// ErrOpt creates the new option with the given value
func ErrOpt(val io.Writer) *Opt {
	return &Opt{Key: ErrOptKey, Val: val}
}

// ErrOptExists determines if the option exists in the given options
func ErrOptExists(opts []*Opt) bool {
	return Exists(opts, ErrOptKey)
}

// GetErrOpt finds and returns the option's value or defaults to os.Stderr
func GetErrOpt(opts []*Opt) io.Writer {
	if o := Get(opts, ErrOptKey); o != nil {
		if val, ok := o.Val.(io.Writer); ok {
			return val
		}
	}
	return os.Stderr
}

// DefaultErrOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultErrOpt(opts []*Opt, val io.Writer) io.Writer {
	if !Exists(opts, ErrOptKey) {
		return val
	}
	return GetErrOpt(opts)
}

// OverwriteErrOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteErrOpt(opts *[]*Opt, val io.Writer) io.Writer {
	return Overwrite(opts, ErrOpt(val)).Val.(io.Writer)
}

// Home Option
// -------------------------------------------------------------------------------------------------

// HomeOpt creates the new option with the given value
func HomeOpt(val string) *Opt {
	return &Opt{Key: HomeOptKey, Val: val}
}

// HomeOptExists determines if the option exists in the given options
func HomeOptExists(opts []*Opt) bool {
	return Exists(opts, HomeOptKey)
}

// GetHomeOpt finds and returns the option's value or defaults to empty string
func GetHomeOpt(opts []*Opt) string {
	if o := Get(opts, HomeOptKey); o != nil {
		if val, ok := o.Val.(string); ok {
			return val
		}
	}
	return ""
}

// DefaultHomeOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultHomeOpt(opts []*Opt, val string) string {
	if !Exists(opts, HomeOptKey) {
		return val
	}
	return GetHomeOpt(opts)
}

// OverwriteHomeOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteHomeOpt(opts *[]*Opt, val string) string {
	return Overwrite(opts, HomeOpt(val)).Val.(string)
}

// Quiet Option
//--------------------------------------------------------------------------------------------------

// QuietOpt creates the new option with the given value
func QuietOpt(val bool) *Opt {
	return &Opt{Key: QuietOptKey, Val: val}
}

// QuietOptExists determines if the option exists in the given options
func QuietOptExists(opts []*Opt) bool {
	return Exists(opts, QuietOptKey)
}

// GetQuietOpt finds and returns the option's value or defaults to false
func GetQuietOpt(opts []*Opt) (result bool) {
	if option := Get(opts, QuietOptKey); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultQuietOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultQuietOpt(opts []*Opt, val bool) bool {
	if !Exists(opts, QuietOptKey) {
		return val
	}
	return GetQuietOpt(opts)
}

// OverwriteQuietOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteQuietOpt(opts *[]*Opt, val bool) bool {
	return Overwrite(opts, QuietOpt(val)).Val.(bool)
}

// Debug Option
//--------------------------------------------------------------------------------------------------

// DebugOpt creates the new option with the given value
func DebugOpt(val bool) *Opt {
	return &Opt{Key: DebugOptKey, Val: val}
}

// DebugOptExists determines if the option exists in the given options
func DebugOptExists(opts []*Opt) bool {
	return Exists(opts, DebugOptKey)
}

// GetDebugOpt finds and returns the option's value or defaults to false
func GetDebugOpt(opts []*Opt) (result bool) {
	if option := Get(opts, DebugOptKey); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultDebugOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultDebugOpt(opts []*Opt, val bool) bool {
	if !Exists(opts, DebugOptKey) {
		return val
	}
	return GetDebugOpt(opts)
}

// OverwriteDebugOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteDebugOpt(opts *[]*Opt, val bool) bool {
	return Overwrite(opts, DebugOpt(val)).Val.(bool)
}

// DryRun Option
//--------------------------------------------------------------------------------------------------

// DryRunOpt creates the new option with the given value
func DryRunOpt(val bool) *Opt {
	return &Opt{Key: DryRunOptKey, Val: val}
}

// DryRunOptExists determines if the option exists in the given options
func DryRunOptExists(opts []*Opt) bool {
	return Exists(opts, DryRunOptKey)
}

// GetDryRunOpt finds and returns the option's value or defaults to false
func GetDryRunOpt(opts []*Opt) (result bool) {
	if option := Get(opts, DryRunOptKey); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultDryRunOpt sets the default value for the option if it doesn't exist already.
// Use this when the Get's default is not desirable.
func DefaultDryRunOpt(opts []*Opt, val bool) bool {
	if !Exists(opts, DryRunOptKey) {
		return val
	}
	return GetDryRunOpt(opts)
}

// OverwriteDryRunOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteDryRunOpt(opts *[]*Opt, val bool) bool {
	return Overwrite(opts, DryRunOpt(val)).Val.(bool)
}

// Testing Option
//--------------------------------------------------------------------------------------------------

// TestingOpt creates the new option with the given value
func TestingOpt(val bool) *Opt {
	return &Opt{Key: TestingOptKey, Val: val}
}

// TestingOptExists determines if the option exists in the given options
func TestingOptExists(opts []*Opt) bool {
	return Exists(opts, TestingOptKey)
}

// GetTestingOpt finds and returns the option's value or defaults to false
func GetTestingOpt(opts []*Opt) (result bool) {
	if option := Get(opts, TestingOptKey); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DefaultTestingOpt returns the option value if found else the one given
// Use this when the Get's default is not desirable.
func DefaultTestingOpt(opts []*Opt, val bool) bool {
	if !Exists(opts, TestingOptKey) {
		return val
	}
	return GetTestingOpt(opts)
}

// OverwriteTestingOpt sets the value for the option to the given value.
// Use this when the new value needs set regardless if the option exists or not.
func OverwriteTestingOpt(opts *[]*Opt, val bool) bool {
	return Overwrite(opts, TestingOpt(val)).Val.(bool)
}
