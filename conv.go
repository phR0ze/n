package n

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Indirect dereferences the interface if needed returning a non-pointer type
func Indirect(obj interface{}) interface{} {
	switch x := obj.(type) {
	case nil:
		return x

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		return x
	case *bool:
		if x == nil {
			return false
		}
		return *x
	case []bool:
		return x
	case []*bool:
		return x
	case *[]bool:
		if x == nil {
			return []bool{}
		}
		return *x
	case *[]*bool:
		if x == nil {
			return []*bool{}
		}
		return *x

	// byte
	//----------------------------------------------------------------------------------------------
	// byte is just a uint8
	// case byte:
	// 	return x
	// case *byte:
	// 	if x == nil {
	// 		return byte(0)
	// 	}
	// 	return *x
	// []byte is just a uint8[] which is defined later
	// case []byte:
	// 	return x
	// case *[]byte:
	// 	if x == nil {
	// 		return []byte{}
	// 	}
	// 	return *x

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		return x
	case *Char:
		if x == nil {
			return *NewChar("")
		}
		return *x
	case []Char:
		return x
	case []*Char:
		return x
	case *[]Char:
		if x == nil {
			return []Char{}
		}
		return *x
	case *[]*Char:
		if x == nil {
			return []*Char{}
		}
		return *x

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		return x
	case *float32:
		if x == nil {
			return float32(0)
		}
		return *x
	case []float32:
		return x
	case []*float32:
		return x
	case *[]float32:
		if x == nil {
			return []float32{}
		}
		return *x
	case *[]*float32:
		if x == nil {
			return []*float32{}
		}
		return *x

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		return x
	case *float64:
		if x == nil {
			return float64(0)
		}
		return *x
	case []float64:
		return x
	case []*float64:
		return x
	case *[]float64:
		if x == nil {
			return []float64{}
		}
		return *x
	case *[]*float64:
		if x == nil {
			return []*float64{}
		}
		return *x

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		return x
	case []*interface{}:
		return x
	case *[]interface{}:
		if x == nil {
			return []interface{}{}
		}
		return *x
	case *[]*interface{}:
		if x == nil {
			return []*interface{}{}
		}
		return *x

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		return x
	case *int:
		if x == nil {
			return 0
		}
		return *x
	case []int:
		return x
	case []*int:
		return x
	case *[]int:
		if x == nil {
			return []int{}
		}
		return *x
	case *[]*int:
		if x == nil {
			return []*int{}
		}
		return *x

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		return x
	case *int8:
		if x == nil {
			return int8(0)
		}
		return *x
	case []int8:
		return x
	case []*int8:
		return x
	case *[]int8:
		if x == nil {
			return []int8{}
		}
		return *x
	case *[]*int8:
		if x == nil {
			return []*int8{}
		}
		return *x

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		return x
	case *int16:
		if x == nil {
			return int16(0)
		}
		return *x
	case []int16:
		return x
	case []*int16:
		return x
	case *[]int16:
		if x == nil {
			return []int16{}
		}
		return *x
	case *[]*int16:
		if x == nil {
			return []*int16{}
		}
		return *x

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		return x
	case *int32:
		if x == nil {
			return int32(0)
		}
		return *x
	case []int32:
		return x
	case []*int32:
		return x
	case *[]int32:
		if x == nil {
			return []int32{}
		}
		return *x
	case *[]*int32:
		if x == nil {
			return []*int32{}
		}
		return *x

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		return x
	case *int64:
		if x == nil {
			return int64(0)
		}
		return *x
	case []int64:
		return x
	case []*int64:
		return x
	case *[]int64:
		if x == nil {
			return []int64{}
		}
		return *x
	case *[]*int64:
		if x == nil {
			return []*int64{}
		}
		return *x

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		return x
	case *IntSlice:
		if x == nil {
			return *NewIntSliceV()
		}
		return *x
	case []IntSlice:
		return x
	case []*IntSlice:
		return x
	case *[]IntSlice:
		if x == nil {
			return []IntSlice{}
		}
		return *x
	case *[]*IntSlice:
		if x == nil {
			return []*IntSlice{}
		}
		return *x

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		return x
	case *Object:
		if x == nil {
			return Object{}
		}
		return *x
	case []Object:
		return x
	case []*Object:
		return x
	case *[]Object:
		if x == nil {
			return []Object{}
		}
		return *x
	case *[]*Object:
		if x == nil {
			return []*Object{}
		}
		return *x

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		return x
	case *Str:
		if x == nil {
			return *NewStrV()
		}
		return *x
	case []Str:
		return x
	case []*Str:
		return x
	case *[]Str:
		if x == nil {
			return []Str{}
		}
		return *x
	case *[]*Str:
		if x == nil {
			return []*Str{}
		}
		return *x

	// rune
	//----------------------------------------------------------------------------------------------
	// rune is a int32 which is already defined
	// case rune:
	// 	return x
	// case *rune:
	// 	if x == nil {
	// 		return rune(0)
	// 	}

	// StringSlice
	//----------------------------------------------------------------------------------------------
	case StringSlice:
		return x
	case *StringSlice:
		if x == nil {
			return *NewStringSliceV()
		}
		return *x
	case []StringSlice:
		return x
	case []*StringSlice:
		return x
	case *[]StringSlice:
		if x == nil {
			return []StringSlice{}
		}
		return *x
	case *[]*StringSlice:
		if x == nil {
			return []*StringSlice{}
		}
		return *x

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		return x
	case *string:
		if x == nil {
			return ""
		}
		return *x
	case []string:
		return x
	case []*string:
		return x
	case *[]string:
		if x == nil {
			return []string{}
		}
		return *x
	case *[]*string:
		if x == nil {
			return []*string{}
		}
		return *x

	// template.CSS
	//----------------------------------------------------------------------------------------------
	case template.CSS:
		return x
	case *template.CSS:
		if x == nil {
			return template.CSS("")
		}
		return *x
	case []template.CSS:
		return x
	case []*template.CSS:
		return x
	case *[]template.CSS:
		if x == nil {
			return []template.CSS{}
		}
		return *x
	case *[]*template.CSS:
		if x == nil {
			return []*template.CSS{}
		}
		return *x

	// template.HTML
	//----------------------------------------------------------------------------------------------
	case template.HTML:
		return x
	case *template.HTML:
		if x == nil {
			return template.HTML("")
		}
		return *x
	case []template.HTML:
		return x
	case []*template.HTML:
		return x
	case *[]template.HTML:
		if x == nil {
			return []template.HTML{}
		}
		return *x
	case *[]*template.HTML:
		if x == nil {
			return []*template.HTML{}
		}
		return *x

	// template.HTMLAttr
	//----------------------------------------------------------------------------------------------
	case template.HTMLAttr:
		return x
	case *template.HTMLAttr:
		if x == nil {
			return template.HTMLAttr("")
		}
		return *x
	case []template.HTMLAttr:
		return x
	case []*template.HTMLAttr:
		return x
	case *[]template.HTMLAttr:
		if x == nil {
			return []template.HTMLAttr{}
		}
		return *x
	case *[]*template.HTMLAttr:
		if x == nil {
			return []*template.HTMLAttr{}
		}
		return *x

	// template.JS
	//----------------------------------------------------------------------------------------------
	case template.JS:
		return x
	case *template.JS:
		if x == nil {
			return template.JS("")
		}
		return *x
	case []template.JS:
		return x
	case []*template.JS:
		return x
	case *[]template.JS:
		if x == nil {
			return []template.JS{}
		}
		return *x
	case *[]*template.JS:
		if x == nil {
			return []*template.JS{}
		}
		return *x

	// template.URL
	//----------------------------------------------------------------------------------------------
	case template.URL:
		return x
	case *template.URL:
		if x == nil {
			return template.URL("")
		}
		return *x
	case []template.URL:
		return x
	case []*template.URL:
		return x
	case *[]template.URL:
		if x == nil {
			return []template.URL{}
		}
		return *x
	case *[]*template.URL:
		if x == nil {
			return []*template.URL{}
		}
		return *x

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		return x
	case *uint:
		if x == nil {
			return uint(0)
		}
		return *x
	case []uint:
		return x
	case []*uint:
		return x
	case *[]uint:
		if x == nil {
			return []uint{}
		}
		return *x
	case *[]*uint:
		if x == nil {
			return []*uint{}
		}
		return *x

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		return x
	case *uint8:
		if x == nil {
			return uint8(0)
		}
		return *x
	case []uint8:
		return x
	case []*uint8:
		return x
	case *[]uint8:
		if x == nil {
			return []uint8{}
		}
		return *x
	case *[]*uint8:
		if x == nil {
			return []*uint8{}
		}
		return *x

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		return x
	case *uint16:
		if x == nil {
			return uint16(0)
		}
		return *x
	case []uint16:
		return x
	case []*uint16:
		return x
	case *[]uint16:
		if x == nil {
			return []uint16{}
		}
		return *x
	case *[]*uint16:
		if x == nil {
			return []*uint16{}
		}
		return *x

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		return x
	case *uint32:
		if x == nil {
			return uint32(0)
		}
		return *x
	case []uint32:
		return x
	case []*uint32:
		return x
	case *[]uint32:
		if x == nil {
			return []uint32{}
		}
		return *x
	case *[]*uint32:
		if x == nil {
			return []*uint32{}
		}
		return *x

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		return x
	case *uint64:
		if x == nil {
			return uint64(0)
		}
		return *x
	case []uint64:
		return x
	case []*uint64:
		return x
	case *[]uint64:
		if x == nil {
			return []uint64{}
		}
		return *x
	case *[]*uint64:
		if x == nil {
			return []*uint64{}
		}
		return *x

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		return reflect.Indirect(reflect.ValueOf(obj)).Interface()
	}
}

// Convert functions
//--------------------------------------------------------------------------------------------------

// ToBool converts an interface to a bool type.
func ToBool(obj interface{}) bool {
	val, _ := ToBoolE(obj)
	return val
}

// ToBoolE converts an interface to a bool type.
func ToBoolE(obj interface{}) (val bool, err error) {
	o := Indirect(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		val = x
	case int:
		val = x != 0
	case int8:
		val = x != 0
	case int16:
		val = x != 0
	case int32:
		val = x != 0
	case int64:
		val = x != 0
	case Object:
		val, err = x.ToBoolE()
	case string:
		if val, err = strconv.ParseBool(x); err != nil {
			err = errors.Wrapf(err, "failed to convert string to bool")
		}
	case uint:
		val = x != 0
	case uint8:
		val = x != 0
	case uint16:
		val = x != 0
	case uint32:
		val = x != 0
	case uint64:
		val = x != 0
	default:
		err = errors.Errorf("unable to convert type %T to bool", x)
	}
	return
}

// ToInt convert an interface to an int type.
func ToInt(obj interface{}) int {
	x, _ := ToIntE(obj)
	return x
}

// ToIntE convert an interface to an int type.
func ToIntE(obj interface{}) (val int, err error) {
	o := Indirect(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToIntE(x.A())
	case float32:
		val = int(x)
	case float64:
		val = int(x)
	case int:
		val = x
	case int8:
		val = int(x)
	case int16:
		val = int(x)
	case int32:
		val = int(x)
	case int64:
		val = int(x)
	case Object:
		val, err = x.ToIntE()
	case uint:
		val = int(x)
	case uint8:
		val = int(x)
	case uint16:
		val = int(x)
	case uint32:
		val = int(x)
	case uint64:
		val = int(x)
	case string:
		var v int64
		if v, err = strconv.ParseInt(x, 0, 0); err != nil {
			err = errors.Wrapf(err, "failed to convert string to int")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = int(v)
		}
	default:
		err = fmt.Errorf("unable to convert type %T to int", x)
	}
	return
}

// ToIntSlice convert an interface to a IntSlice type which will never be nil
func ToIntSlice(obj interface{}) *IntSlice {
	x, _ := ToIntSliceE(obj)
	if x == nil {
		return &IntSlice{}
	}
	return x
}

// ToIntSliceE convert an interface to a IntSlice type.
func ToIntSliceE(obj interface{}) (val *IntSlice, err error) {
	val = &IntSlice{}
	o := Indirect(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		*val = append(*val, ToInt(x))
	case []bool:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}
	case []*bool:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		*val = append(*val, ToInt(x))
	case []Char:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}
	case []*Char:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		*val = append(*val, int(x))
	case []float32:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*float32:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		*val = append(*val, int(x))
	case []float64:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*float64:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.WithMessagef(e, "unable to convert %T to []int", x)
			}
		}

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		*val = append(*val, x)
	case []int:
		v := IntSlice(x)
		val = &v
	case []*int:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		*val = append(*val, int(x))
	case []int8:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int8:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		*val = append(*val, int(x))
	case []int16:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int16:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		*val = append(*val, int(x))
	case []int32:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int32:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		*val = append(*val, int(x))
	case []int64:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int64:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		val = &x
	case []IntSlice:
		for i := range x {
			*val = append(*val, x[i]...)
		}
	case []*IntSlice:
		for i := range x {
			if x[i] != nil {
				*val = append(*val, *x[i]...)
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		if v, e := x.ToIntE(); e != nil {
			err = fmt.Errorf("unable to convert type Object to []int")
		} else {
			*val = append(*val, v)
		}
	case []Object:
		for i := range x {
			if v, e := x[i].ToIntE(); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type %T to []int", x[i].O())
				return
			}
		}
	case []*Object:
		for i := range x {
			if v, e := x[i].ToIntE(); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type %T to []int", x[i].O())
				return
			}
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		if v, e := ToIntE(x); e != nil {
			err = fmt.Errorf("unable to convert type string to []int")
		} else {
			*val = append(*val, v)
		}
	case []string:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type string to []int")
				return
			}
		}
	case []*string:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type *string to []int")
				return
			}
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		*val = append(*val, int(x))
	case []uint:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		*val = append(*val, int(x))
	case []uint8:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint8:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		*val = append(*val, int(x))
	case []uint16:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint16:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		*val = append(*val, int(x))
	case []uint32:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint32:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		*val = append(*val, int(x))

	case []uint64:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint64:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(x)
		k := v.Kind()

		switch {

		// generically convert array and slice types
		case k == reflect.Array || k == reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				if v, e := ToIntE(v.Index(i).Interface()); e == nil {
					*val = append(*val, v)
				} else {
					err = errors.WithMessagef(e, "unable to convert %T to []int", x)
					return
				}
			}

		// not supporting this type yet
		default:
			err = fmt.Errorf("unable to convert type %T to []int", x)
			return
		}
	}
	return
}

// ToString convert an interface to a string type.
func ToString(obj interface{}) string {
	o := Indirect(obj)

	switch x := o.(type) {
	case nil:
		return ""
	case string:
		return x
	case bool:
		return strconv.FormatBool(x)
	case []byte:
		return string(x)
	case Char:
		return x.String()
	case error:
		return x.Error()
	case float32:
		return strconv.FormatFloat(float64(x), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(x), 10)
	case int8:
		return strconv.FormatInt(int64(x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case int64:
		return strconv.FormatInt(x, 10)
	case template.CSS:
		return string(x)
	case template.HTML:
		return string(x)
	case template.HTMLAttr:
		return string(x)
	case template.JS:
		return string(x)
	case template.URL:
		return string(x)
	case []rune:
		return string(x)
	case uint:
		return strconv.FormatInt(int64(x), 10)
	case uint8:
		return strconv.FormatInt(int64(x), 10)
	case uint16:
		return strconv.FormatInt(int64(x), 10)
	case uint32:
		return strconv.FormatInt(int64(x), 10)
	case uint64:
		return strconv.FormatInt(int64(x), 10)
	case fmt.Stringer, *fmt.Stringer:
		if x, ok := x.(*fmt.Stringer); ok {
			if x == nil {
				return ""
			}
			return (*x).String()
		}
		return x.(fmt.Stringer).String()
	default:
		return fmt.Sprintf("%v", obj)
	}
}

// ToStringSlice convert an interface to a []int type.
func ToStringSlice(obj interface{}) []string {
	x, _ := ToStringSliceE(obj)
	if x == nil {
		return []string{}
	}
	return x
}

// ToStringSliceE convert an interface to a []string type.
func ToStringSliceE(obj interface{}) (val []string, err error) {
	val = []string{}
	o := Indirect(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	// case bool:
	// 	val = []int{ToInt(x)}
	// case []bool:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}
	// case []*bool:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // Char
	// //----------------------------------------------------------------------------------------------
	// case Char:
	// 	val = []int{ToInt(x)}
	// case []Char:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}
	// case []*Char:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // float32
	// //----------------------------------------------------------------------------------------------
	// case float32:
	// 	val = []int{int(x)}
	// case []float32:
	// 	for i := range x {
	// 		val = append(val, int(x[i]))
	// 	}
	// case []*float32:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // float64
	// //----------------------------------------------------------------------------------------------
	// case float64:
	// 	val = []int{int(x)}
	// case []float64:
	// 	for i := range x {
	// 		val = append(val, int(x[i]))
	// 	}
	// case []*float64:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// // int
	// //----------------------------------------------------------------------------------------------
	// case int:
	// 	val = []int{int(x)}
	// case []int:
	// 	val = x
	// case []*int:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // int8
	// //----------------------------------------------------------------------------------------------
	// case int8:
	// 	val = []int{int(x)}
	// case []int8:
	// 	for i := range x {
	// 		val = append(val, int(x[i]))
	// 	}
	// case []*int8:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // int16
	// //----------------------------------------------------------------------------------------------
	// case int16:
	// 	val = []int{int(x)}
	// case []int16:
	// 	for i := range x {
	// 		val = append(val, int(x[i]))
	// 	}
	// case []*int16:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // int32
	// //----------------------------------------------------------------------------------------------
	// case int32:
	// 	val = []int{int(x)}
	// case []int32:
	// 	for i := range x {
	// 		val = append(val, int(x[i]))
	// 	}
	// case []*int32:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // int64
	// //----------------------------------------------------------------------------------------------
	// case int64:
	// 	val = []int{int(x)}
	// case []int64:
	// 	for i := range x {
	// 		val = append(val, int(x[i]))
	// 	}
	// case []*int64:
	// 	for i := range x {
	// 		val = append(val, ToInt(x[i]))
	// 	}

	// // IntSlice
	// //----------------------------------------------------------------------------------------------
	// case IntSlice:
	// 	val = x
	// case []IntSlice:
	// 	for i := range x {
	// 		val = append(val, x[i]...)
	// 	}
	// case []*IntSlice:
	// 	for i := range x {
	// 		if x[i] != nil {
	// 			val = append(val, *x[i]...)
	// 		}
	// 	}

	// // Object
	// //----------------------------------------------------------------------------------------------
	// case Object:
	// 	if v, e := x.ToIntE(); e != nil {
	// 		err = fmt.Errorf("unable to convert type Object to []int")
	// 	} else {
	// 		val = []int{v}
	// 	}
	// case []Object:
	// 	for i := range x {
	// 		if v, e := x[i].ToIntE(); e == nil {
	// 			val = append(val, v)
	// 		} else {
	// 			err = fmt.Errorf("unable to convert type %T to []int", x[i].O())
	// 			return
	// 		}
	// 	}
	// case []*Object:
	// 	for i := range x {
	// 		if v, e := x[i].ToIntE(); e == nil {
	// 			val = append(val, v)
	// 		} else {
	// 			err = fmt.Errorf("unable to convert type %T to []int", x[i].O())
	// 			return
	// 		}
	// 	}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		val = strings.Fields(x)
	case []string:
		val = x
	case []*string:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

		// // uint
		// //----------------------------------------------------------------------------------------------
		// case uint:
		// 	val = []int{int(x)}
		// case []uint:
		// 	for i := range x {
		// 		val = append(val, int(x[i]))
		// 	}
		// case []*uint:
		// 	for i := range x {
		// 		val = append(val, ToInt(x[i]))
		// 	}

		// // uint8
		// //----------------------------------------------------------------------------------------------
		// case uint8:
		// 	val = []int{int(x)}
		// case []uint8:
		// 	for i := range x {
		// 		val = append(val, int(x[i]))
		// 	}
		// case []*uint8:
		// 	for i := range x {
		// 		val = append(val, ToInt(x[i]))
		// 	}

		// // uint16
		// //----------------------------------------------------------------------------------------------
		// case uint16:
		// 	val = []int{int(x)}
		// case []uint16:
		// 	for i := range x {
		// 		val = append(val, int(x[i]))
		// 	}
		// case []*uint16:
		// 	for i := range x {
		// 		val = append(val, ToInt(x[i]))
		// 	}

		// // uint32
		// //----------------------------------------------------------------------------------------------
		// case uint32:
		// 	val = []int{int(x)}
		// case []uint32:
		// 	for i := range x {
		// 		val = append(val, int(x[i]))
		// 	}
		// case []*uint32:
		// 	for i := range x {
		// 		val = append(val, ToInt(x[i]))
		// 	}

		// // uint64
		// //----------------------------------------------------------------------------------------------
		// case uint64:
		// 	val = []int{int(x)}

		// case []uint64:
		// 	for i := range x {
		// 		val = append(val, int(x[i]))
		// 	}
		// case []*uint64:
		// 	for i := range x {
		// 		val = append(val, ToInt(x[i]))
		// 	}

		// // fall back on reflection
		// //----------------------------------------------------------------------------------------------
		// default:
		// 	v := reflect.ValueOf(x)
		// 	k := v.Kind()

		// 	switch {

		// 	// generically convert array and slice types
		// 	case k == reflect.Array || k == reflect.Slice:
		// 		for i := 0; i < v.Len(); i++ {
		// 			if v, e := ToIntE(v.Index(i).Interface()); e == nil {
		// 				val = append(val, v)
		// 			} else {
		// 				err = errors.WithMessagef(e, "unable to convert %T to []int", x)
		// 				return
		// 			}
		// 		}

		// 	// not supporting this type yet
		// 	default:
		// 		err = fmt.Errorf("unable to convert type %T to []int", x)
		// 		return
		// 	}
	}
	return
}
