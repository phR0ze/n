// Package structs provides reflection based utilities for working with structs.
//
// structs was primarily created to provide a way to initialize a struct of string
// properties with the values of the property names. In this way we can mimic the
// ability to initialize nested consts available in other languages. It is based
// on reflection so does have a performance penalty however it is meant to be
// used as a global var or initialized with init for a one time cost
package structs

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Struct provides a reflection wrapper for structs
type Struct struct {
	o  interface{}   // struct object were working with
	v  reflect.Value // value to object
	ov reflect.Value // value of original object
	k  reflect.Kind  // kind of the object
	ok reflect.Kind  // kind of original object
}

// New converts a given struct to a *Struct to work with it in reflection;
// panics if obj is not a struct type.
func New(obj interface{}) (new *Struct) {
	new = &Struct{o: obj}

	// Get reflection handles
	new.ov = reflect.ValueOf(obj)
	new.ok = new.ov.Kind()

	// Dereference pointers
	if new.ok == reflect.Ptr {
		new.v = new.ov.Elem()
		new.k = new.v.Kind()
	} else {
		new.v = new.ov
		new.k = new.ok
	}

	// Panic if we don't have the right kind
	if new.k != reflect.Struct {
		panic(fmt.Sprintf("structs.New requires a non nil struct type not a %v", new.k))
	}

	return
}

// Init all the string fields to the names of the fields;
// panics if obj is not a struct pointer type or is nil.
func Init(obj interface{}) {
	st := New(obj)
	if st.ok != reflect.Ptr {
		panic(fmt.Sprintf("structs.InitStrs requires a struct pointer type not a %v", st.ok))
	}

	strtype := reflect.TypeOf("")
	for i, field := range st.Fields() {
		if field.Type == strtype {
			vfield := st.v.Field(i)

			// Get the masked value of private fields via its address
			if field.PkgPath != "" {
				vfield = reflect.NewAt(field.Type, unsafe.Pointer(vfield.UnsafeAddr())).Elem()
			}

			vfield.SetString(field.Name)
		}
	}
}

// Fields returns all the field names as a []string
func (st *Struct) Fields() (fields []reflect.StructField) {
	typ := st.v.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fields = append(fields, field)
	}
	return
}

// SetField sets the value of the field to the given value if they are the same type
func (st *Struct) SetField() {

}
