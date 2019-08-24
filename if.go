package n

import (
	"reflect"

	"github.com/pkg/errors"
)

// IF provides a way to deal with conditionals more elegantly in Go
type IF struct {
	State  bool          // Status of the result
	Error  error         // Errors that may have been captured
	Return []interface{} // Return values
}

// If provides a way to execute one line conditionals in Go
func If(state bool, err ...error) *IF {
	cond := &IF{State: state}
	if len(err) > 0 {
		cond.Error = err[0]
	}
	return cond
}

// Do executes the given function f with the given paramaters p if IF.State is true.
// When IF.State is true the IF object is Reset and the properties are used for the
// state of the current Do execution returning all return values from the function
// given to Do in the IF.Return slice and the first identified return bool will be
// used for IF.State and the first identified return error will be used for IF.Error.
func (cond *IF) Do(f interface{}, params ...interface{}) *IF {
	if !cond.State {
		return cond
	}
	cond.Reset()
	vf := reflect.ValueOf(f)

	// Ensure the target function is of the correctd type and params
	if vf.Kind() != reflect.Func {
		cond.Error = errors.Errorf("target function is not of type reflect.Func")
		return cond
	}
	if vf.Type().NumIn() != len(params) {
		cond.Error = errors.Errorf("incorrect number of parameters for the given function")
		return cond
	}

	// Convert the given params into a slice of reflect.Value
	vp := []reflect.Value{}
	for _, param := range params {
		vp = append(vp, reflect.ValueOf(param))
	}

	// Execute the target function
	errSet, stateSet := false, false
	for _, val := range vf.Call(vp) {
		obj := val.Interface()
		cond.Return = append(cond.Return, obj)

		if !stateSet || !errSet {
			switch x := obj.(type) {
			case bool:
				if !stateSet {
					cond.State = x
					stateSet = true
				}
			case error:
				if !errSet {
					cond.Error = x
					errSet = true
				}
			}
		}
	}

	return cond
}

// Reset defaults all the IF properties
func (cond *IF) Reset() *IF {
	cond.Error = nil
	cond.State = false
	cond.Return = []interface{}{}
	return cond
}
