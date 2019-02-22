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
	for _, opt := range opts {
		if opt.Key == key {
			return opt
		}
	}
	return nil
}
