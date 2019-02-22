// Package opt provides a simple struct for options that can be passed
// in as an optional last variadic parameter.
package opt

// Opt provides a mechanism for passing in optional options to functions.
type Opt struct {
	Key string      // name of the option
	Val interface{} // value of the option
}

// Find an option by the given key
func Find(opts []*Opt, key string) *Opt {
	for _, o := range opts {
		if o.Key == key {
			return o
		}
	}
	return nil
}

// Options implementation pattern.
// Create helper functions that wrap your constants to allow the compiler to assist.
//--------------------------------------------------------------------------------------------------

// DefaultDebugOpt sets the default for the flag
func DefaultDebugOpt(opts *[]*Opt, val bool) {
	if option := Find(*opts, "debug"); option == nil {
		*opts = append(*opts, &Opt{Key: "debug", Val: val})
	}
}

// DefaultDryRunOpt sets the default for the flag
func DefaultDryRunOpt(opts *[]*Opt, val bool) {
	if option := Find(*opts, "dry-run"); option == nil {
		*opts = append(*opts, &Opt{Key: "dry-run", Val: val})
	}
}

// DebugOpt finds and returns the option's value or a default value if missing
func DebugOpt(opts []*Opt) (result bool) {
	if option := Find(opts, "debug"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DryRunOpt finds and returns the option's value or a default value if missing
func DryRunOpt(opts []*Opt) (result bool) {
	if option := Find(opts, "dry-run"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// NewDebugOpt creates the new option with the given value
func NewDebugOpt(val bool) *Opt {
	return &Opt{Key: "debug", Val: val}
}

// NewDryRunOpt creates the new option with the given value
func NewDryRunOpt(val bool) *Opt {
	return &Opt{Key: "dry-run", Val: val}
}
