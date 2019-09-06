// Package opt provides a simple struct for options that can be passed
// in as an optional last variadic parameter. Packages making use of this
// pattern should ignore options that are not supported and clearly call out
// in documentation which options are supported. Please provide local option
// help function implementations for options rather than using magic strings.
package opt

// Opt provides a mechanism for passing in optional options to functions.
type Opt struct {
	Key string      // name of the option
	Val interface{} // value of the option
}

// Find an option by the given key
func Find(opts []*Opt, key string) *Opt {
	for _, o := range opts {
		if o != nil && o.Key == key {
			return o
		}
	}
	return nil
}

// Options implementation pattern.
// Create helper functions that wrap your constants to allow the compiler to assist.
// The creation functions should be called <option>Opt and the getters Get<option>Opt
// and the add defeaults Default<option>Opt
//--------------------------------------------------------------------------------------------------

// DefaultDebugOpt sets the default for the flag if it doesn't exist already
func DefaultDebugOpt(opts *[]*Opt, val bool) {
	if option := Find(*opts, "debug"); option == nil {
		*opts = append(*opts, &Opt{Key: "debug", Val: val})
	}
}

// DefaultDryRunOpt sets the default for the flag if it doesn't exist already
func DefaultDryRunOpt(opts *[]*Opt, val bool) {
	if option := Find(*opts, "dry-run"); option == nil {
		*opts = append(*opts, &Opt{Key: "dry-run", Val: val})
	}
}

// GetDebugOpt finds and returns the option's value or a default value if missing
func GetDebugOpt(opts []*Opt) (result bool) {
	if option := Find(opts, "debug"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// GetDryRunOpt finds and returns the option's value or a default value if missing
func GetDryRunOpt(opts []*Opt) (result bool) {
	if option := Find(opts, "dry-run"); option != nil {
		if val, ok := option.Val.(bool); ok {
			result = val
		}
	}
	return
}

// DebugOpt creates the new option with the given value
func DebugOpt(val bool) *Opt {
	return &Opt{Key: "debug", Val: val}
}

// DryRunOpt creates the new option with the given value
func DryRunOpt(val bool) *Opt {
	return &Opt{Key: "dry-run", Val: val}
}
