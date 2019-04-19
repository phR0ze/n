package n

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// Indirect dereferences the interface if needed returning a non-pointer type
func Indirect(obj interface{}) interface{} {
	switch x := obj.(type) {
	case nil:
		return x
	case bool:
		return x
	case *bool:
		if x == nil {
			return false
		}
		return *x
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
	case Char:
		return x
	case *Char:
		if x == nil {
			return *NewChar("")
		}
		return *x
	case float32:
		return x
	case *float32:
		if x == nil {
			return float32(0)
		}
		return *x
	case float64:
		return x
	case *float64:
		if x == nil {
			return float64(0)
		}
		return *x
	case int:
		return x
	case *int:
		if x == nil {
			return 0
		}
		return *x
	case int8:
		return x
	case *int8:
		if x == nil {
			return int8(0)
		}
		return *x
	case int16:
		return x
	case *int16:
		if x == nil {
			return int16(0)
		}
		return *x
	case int32:
		return x
	case *int32:
		if x == nil {
			return int32(0)
		}
		return *x
	case int64:
		return x
	case *int64:
		if x == nil {
			return int64(0)
		}
		return *x
	case []int:
		return x
	case *[]int:
		if x == nil {
			return []int{}
		}
		return *x
	case *[]int8:
		if x == nil {
			return []int8{}
		}
		return *x
	case []int16:
		return x
	case *[]int16:
		if x == nil {
			return []int16{}
		}
		return *x
	case []int32:
		return x
	case *[]int32:
		if x == nil {
			return []int32{}
		}
		return *x
	case []int64:
		return x
	case *[]int64:
		if x == nil {
			return []int64{}
		}
		return *x
	case IntSlice:
		return x
	case *IntSlice:
		if x == nil {
			return *NewIntSliceV()
		}
		return *x
	case Str:
		return x
	case *Str:
		if x == nil {
			return *NewStrV()
		}
		return *x
	// rune is a int32 which is already defined
	// case rune:
	// 	return x
	// case *rune:
	// 	if x == nil {
	// 		return rune(0)
	// 	}
	case StringSlice:
		return x
	case *StringSlice:
		if x == nil {
			return *NewStringSliceV()
		}
		return *x
	case string:
		return x
	case *string:
		if x == nil {
			return ""
		}
		return *x
	case []string:
		return x
	case *[]string:
		if x == nil {
			return []string{}
		}
		return *x
	case template.CSS:
		return x
	case *template.CSS:
		if x == nil {
			return template.CSS("")
		}
		return *x
	case template.HTML:
		return x
	case *template.HTML:
		if x == nil {
			return template.HTML("")
		}
		return *x
	case template.HTMLAttr:
		return x
	case *template.HTMLAttr:
		if x == nil {
			return template.HTMLAttr("")
		}
		return *x
	case template.JS:
		return x
	case *template.JS:
		if x == nil {
			return template.JS("")
		}
		return *x
	case template.URL:
		return x
	case *template.URL:
		if x == nil {
			return template.URL("")
		}
		return *x
	case uint:
		return x
	case *uint:
		if x == nil {
			return uint(0)
		}
		return *x
	case uint8:
		return x
	case *uint8:
		if x == nil {
			return uint8(0)
		}
		return *x
	case uint16:
		return x
	case *uint16:
		if x == nil {
			return uint16(0)
		}
		return *x
	case uint32:
		return x
	case *uint32:
		if x == nil {
			return uint32(0)
		}
		return *x
	case uint64:
		return x
	case *uint64:
		if x == nil {
			return uint64(0)
		}
		return *x
	case []uint:
		return x
	case *[]uint:
		if x == nil {
			return []uint{}
		}
		return *x
	case []uint8:
		return x
	case *[]uint8:
		if x == nil {
			return []uint8{}
		}
		return *x
	case []uint16:
		return x
	case *[]uint16:
		if x == nil {
			return []uint16{}
		}
		return *x
	case []uint32:
		return x
	case *[]uint32:
		if x == nil {
			return []uint32{}
		}
		return *x
	case []uint64:
		return x
	case *[]uint64:
		if x == nil {
			return []uint64{}
		}
		return *x
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
		err = errors.Errorf("unable to cast type %T to bool", x)
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
	case uint:
		val = int(x)
	case uint64:
		val = int(x)
	case uint32:
		val = int(x)
	case uint16:
		val = int(x)
	case uint8:
		val = int(x)
	case float32:
		val = int(x)
	case float64:
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
		err = fmt.Errorf("unable to cast type %T to int", x)
	}
	return
}

// ToIntSlice convert an interface to a []int type.
func ToIntSlice(obj interface{}) []int {
	x, _ := ToIntSliceE(obj)
	if x == nil {
		return []int{}
	}
	return x
}

// ToIntSliceE convert an interface to a []int type.
func ToIntSliceE(obj interface{}) (val []int, err error) {
	o := Indirect(obj)

	switch x := o.(type) {
	case nil:
		val = []int{}
	case []int:
		val = x
	case IntSlice:
		val = x
	}

	// kind := reflect.TypeOf(i).Kind()
	// switch kind {
	// case reflect.Slice, reflect.Array:
	// 	s := reflect.ValueOf(i)
	// 	a := make([]int, s.Len())
	// 	for j := 0; j < s.Len(); j++ {
	// 		val, err := ToIntE(s.Index(j).Interface())
	// 		if err != nil {
	// 			return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	// 		}
	// 		a[j] = val
	// 	}
	// 	return a, nil
	// default:
	// 	return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	// }
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
