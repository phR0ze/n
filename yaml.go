package n

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Yaml gets data by key which can be dot delimited
// returns nil queryable on errors or keys not found
func (q *Queryable) Yaml(key string) (result *Queryable) {
	keys := A(key).Split(".")
	if key, ok := keys.TakeFirst(); ok {
		switch x := q.v.Interface().(type) {
		case map[string]interface{}:
			if !A(key).ContainsAny(":", "[", "]") {
				if v, ok := x[key]; ok {
					result = Q(v)
				}
			}
		case []interface{}:
			k, v := A(key).TrimPrefix("[").TrimSuffix("]").Split(":").YamlPair()
			if v == nil {
				if i, err := strconv.Atoi(k); err == nil {
					result = q.At(i)
				} else {
					panic(errors.New("Failed to convert index to an int"))
				}
			} else {
				for i := range x {
					if m, ok := x[i].(map[string]interface{}); ok {
						if entry, ok := m[k]; ok {
							if v == entry {
								result = Q(m)
								break
							}
						}
					}
				}
			}
		}
		if keys.Len() != 0 && result != nil && result.Any() {
			result = result.Yaml(keys.Join(".").A())
		}
	}
	if result == nil {
		result = N()
	}
	return
}

// YamlReplace recursively makes string substitutions
func YamlReplace(data interface{}, values map[string]string) (result interface{}) {
	switch x := data.(type) {
	case map[string]interface{}:
		for k, v := range x {
			switch y := v.(type) {
			case string:
				for o, n := range values {
					y = strings.Replace(y, o, n, -1)
				}
				x[k] = y
			case []interface{}, map[string]interface{}:
				x[k] = YamlReplace(v, values)
			}
		}
		result = data
	case []map[string]interface{}:
		resultSlice := []map[string]interface{}{}
		for i := range x {
			v := YamlReplace(x[i], values)
			resultSlice = append(resultSlice, v.(map[string]interface{}))
		}
		result = resultSlice
	case []interface{}:
		resultSlice := []interface{}{}
		for i := range x {
			var value interface{}
			switch y := x[i].(type) {
			case string:
				for o, n := range values {
					y = strings.Replace(y, o, n, -1)
				}
				value = y
			case []interface{}, map[string]interface{}:
				value = YamlReplace(y, values)
			default:
				value = y
			}
			resultSlice = append(resultSlice, value)
		}
		result = resultSlice
	default:
		result = data
	}
	return
}

// YamlSet sets data by key which can be dot delimited
func (q *Queryable) YamlSet(key string, data interface{}) (result *Queryable, err error) {
	keys := A(key).Split(".")
	if key, ok := keys.TakeFirst(); ok {
		switch x := q.v.Interface().(type) {
		case map[string]interface{}:

			// Current target is a map key
			if !A(key).ContainsAny(":", "[", "]") {

				// No more keys so we've reached our destination
				if !keys.Any() {
					x[key] = data
				} else {
					var v interface{}
					if v, ok = x[key]; !ok {
						// Doesn't exist so create
						x[key] = map[string]interface{}{}
						v = x[key]
					}
					result, err = Q(v).YamlSet(keys.Join(".").A(), data)
				}
			}
		case []interface{}:
			k, v := A(key).TrimPrefix("[").TrimSuffix("]").Split(":").YamlPair()
			if v == nil {
				var i int
				if i, err = strconv.Atoi(k); err == nil {

					// No more keys so we've reached our destination
					if !keys.Any() {
						if q.Len() > i {
							// Override
							//recurse = q.At(i)
						} else {
							// Insert new item
							q.Append(data)
						}
					} else {
						if i < q.Len() {
							result, err = q.At(i).YamlSet(keys.Join(".").A(), data)
						} else {
							err = fmt.Errorf("Indexing out of bounds")
							return
						}
					}
				} else {
					return
				}
			} else {
				for i := range x {
					if m, ok := x[i].(map[string]interface{}); ok {
						if entry, ok := m[k]; ok {
							if v == entry {
								if !keys.Any() {
									//Q(m)
								} else {
									result, err = q.At(i).YamlSet(keys.Join(".").A(), data)
								}
								break
							}
						}
					}
				}
			}
		}
	}
	result = q
	return
}

// Insert/set data in the unmarshalled yaml
func (q *Queryable) yamlSet() (err error) {
	return
}
