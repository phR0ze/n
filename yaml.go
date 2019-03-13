package n

// // Yaml gets data by key which can be dot delimited
// // returns nil queryable on errors or keys not found
// func (q *Queryable) Yaml(key string) (result *Queryable) {
// 	keys := A(key).Split(".")
// 	if key, ok := keys.TakeFirst(); ok {
// 		switch x := q.v.Interface().(type) {
// 		case map[string]interface{}:
// 			if !A(key).ContainsAny(":", "[", "]") {
// 				if v, ok := x[key]; ok {
// 					result = Q(v)
// 				}
// 			}
// 		case []interface{}:
// 			k, v := A(key).TrimPrefix("[").TrimSuffix("]").Split(":").YamlPair()
// 			if v == nil {
// 				if i, err := strconv.Atoi(k); err == nil {
// 					result = q.At(i)
// 				} else {
// 					panic(errors.New("Failed to convert index to an int"))
// 				}
// 			} else {
// 				for i := range x {
// 					if m, ok := x[i].(map[string]interface{}); ok {
// 						if entry, ok := m[k]; ok {
// 							if v == entry {
// 								result = Q(m)
// 								break
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 		if keys.Len() != 0 && result != nil && result.Any() {
// 			result = result.Yaml(keys.Join(".").A())
// 		}
// 	}
// 	if result == nil {
// 		result = N()
// 	}
// 	return
// }

// // YamlMerge merges the given yaml into the existing yaml
// // with the given yaml having a higher precedence than the existing
// func (q *Queryable) YamlMerge(data interface{}) (result *Queryable, err error) {
// 	result = q
// 	switch x := q.v.Interface().(type) {
// 	case map[string]interface{}:
// 		switch y := data.(type) {
// 		case map[string]interface{}:
// 			q.Copy(MergeMap(x, y))
// 		default:
// 			err = fmt.Errorf("Invalid merge type")
// 		}
// 	case []interface{}:
// 		switch y := data.(type) {
// 		case []interface{}:
// 			for i := range y {
// 				if i < len(x) {
// 					x[i] = y[i]
// 				} else {
// 					x = append(x, y[i])
// 					q.Copy(x)
// 				}
// 			}
// 		default:
// 			err = fmt.Errorf("Invalid merge type")
// 		}
// 	}
// 	return
// }

// // YamlSet sets data by key which can be dot delimited
// func (q *Queryable) YamlSet(key string, data interface{}) (result *Queryable, err error) {
// 	keys := A(key).Split(".")
// 	if key, ok := keys.TakeFirst(); ok {
// 		switch x := q.v.Interface().(type) {
// 		case map[string]interface{}:

// 			// Current target is a map key
// 			if !A(key).ContainsAny(":", "[", "]") {

// 				// No more keys so we've reached our destination
// 				if !keys.Any() {
// 					x[key] = data
// 				} else {
// 					var v interface{}
// 					if v, ok = x[key]; !ok {
// 						// Doesn't exist so create
// 						if keys.First().Contains("[") {
// 							x[key] = []interface{}{}
// 						} else {
// 							x[key] = map[string]interface{}{}
// 						}
// 						v = x[key]
// 					}
// 					if result, err = Q(v).YamlSet(keys.Join(".").A(), data); err == nil {
// 						x[key] = result.O()
// 					}
// 				}
// 			}
// 		case []interface{}:
// 			insert := false
// 			var k string
// 			var v interface{}
// 			if A(key).Contains("[[") {
// 				insert = true
// 				k, v = A(key).TrimPrefix("[[").TrimSuffix("]]").Split(":").YamlPair()
// 			} else {
// 				k, v = A(key).TrimPrefix("[").TrimSuffix("]").Split(":").YamlPair()
// 			}
// 			if v == nil {
// 				var i int
// 				if i, err = strconv.Atoi(k); err == nil {

// 					// No more keys so we've reached our destination
// 					if !keys.Any() {
// 						if i < q.Len() && !insert {
// 							q.Set(i, data)
// 						} else {
// 							q.Insert(i, data)
// 						}
// 					} else {
// 						if i >= q.Len() {
// 							if keys.First().Contains("[") {
// 								q.Append([]interface{}{})
// 							} else {
// 								q.Append(map[string]interface{}{})
// 							}
// 						}
// 						result, err = q.At(i).YamlSet(keys.Join(".").A(), data)
// 					}
// 				} else {
// 					return
// 				}
// 			} else {
// 				for i := range x {
// 					if m, ok := x[i].(map[string]interface{}); ok {
// 						if entry, ok := m[k]; ok {
// 							if v == entry {
// 								if !keys.Any() {
// 									if insert {
// 										q.Insert(i, data)
// 									} else {
// 										q.Set(i, data)
// 									}
// 								} else {
// 									result, err = q.At(i).YamlSet(keys.Join(".").A(), data)
// 								}
// 								break
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	result = q
// 	return
// }

// // Insert/set data in the unmarshalled yaml
// func (q *Queryable) yamlSet() (err error) {
// 	return
// }
