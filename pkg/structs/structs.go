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
	v reflect.Value // value to object
	k reflect.Kind  // kind of the object
	p reflect.Value // pointer to object
}

// New converts a given struct to a *Struct to work with it in reflection;
// panics if obj is not a non nil struct/*struct type.
func New(obj interface{}) (new *Struct) {
	new = &Struct{}

	// Ensure we always have an addressible type
	new.v = reflect.ValueOf(obj)
	if new.v.Kind() != reflect.Ptr {
		v := reflect.New(new.v.Type())
		v.Elem().Set(new.v)
		new.v = v
	}
	new.p = new.v
	new.v = new.v.Elem()
	new.k = new.v.Kind()

	// Panic if we don't have the right kind
	if new.k != reflect.Struct {
		panic(fmt.Sprintf("structs.New requires a non nil struct type not a %v", new.k))
	}

	return
}

// Init all the string fields to the names of the fields;
// panics if obj is not a non nil struct/*struct type.
func Init(obj interface{}) {
	st := New(obj)
	strtype := reflect.TypeOf("")
	for i, field := range st.Fields() {
		if field.Type == strtype {
			st.SetFieldByIndex(i, field.Name, strtype, field.Type)
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

// SetFieldByIndex sets the value of the field to the given value if they are the same type;
// Sets private or public fields as called out by the index i. Optionially takes the value's
// type and the field's value.
func (st *Struct) SetFieldByIndex(i int, value interface{}, typ ...reflect.Type) {
	vfield := st.v.Field(i)

	// Determine value type
	var valueType reflect.Type
	if len(typ) != 0 {
		valueType = typ[0]
	} else {
		valueType = reflect.TypeOf(value)
	}

	// Determine the field type
	var fieldType reflect.Type
	if len(typ) > 1 {
		fieldType = typ[1]
	} else {
		fieldType = vfield.Type()
	}

	// Ensure we have the same types
	if fieldType == valueType {

		// Get the masked value of private fields via its address
		if !vfield.CanSet() {
			vfield = reflect.NewAt(fieldType, unsafe.Pointer(vfield.UnsafeAddr())).Elem()
		}

		// Set the value according to type
		switch x := value.(type) {
		case bool:
			vfield.SetBool(x)
		case float32:
			vfield.SetFloat(float64(x))
		case float64:
			vfield.SetFloat(x)
		case int:
			vfield.SetInt(int64(x))
		case int8:
			vfield.SetInt(int64(x))
		case int16:
			vfield.SetInt(int64(x))
		case int32:
			vfield.SetInt(int64(x))
		case int64:
			vfield.SetInt(x)
		case uint:
			vfield.SetUint(uint64(x))
		case uint8:
			vfield.SetUint(uint64(x))
		case uint16:
			vfield.SetUint(uint64(x))
		case uint32:
			vfield.SetUint(uint64(x))
		case uint64:
			vfield.SetUint(x)
		case string:
			vfield.SetString(x)
		}
	}
}
