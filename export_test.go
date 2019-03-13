package n

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestA(t *testing.T) {
// 	{
// 		q := Q("one")
// 		assert.Equal(t, "o", q.At(0).A())
// 	}
// 	{
// 		q := Q([]string{"one"})
// 		assert.Equal(t, "one", q.At(0).A())
// 		assert.Equal(t, []string{"one"}, q.Strs())
// 	}
// 	{
// 		assert.Equal(t, "1", Q("1").A())
// 		assert.Equal(t, "1", Q(1).A())
// 		assert.Equal(t, "true", Q(true).A())
// 		assert.Equal(t, "false", Q(false).A())
// 	}
// }

// func TestB(t *testing.T) {
// 	{
// 		assert.Equal(t, true, getB(t, Q(true)))
// 		assert.Equal(t, false, getB(t, Q(false)))
// 	}
// 	{
// 		q := Q([]bool{true})
// 		assert.Equal(t, true, getB(t, q.At(0)))
// 	}
// 	{
// 		assert.Equal(t, true, getB(t, Q(4)))
// 		assert.Equal(t, false, getB(t, Q(0)))
// 	}
// 	{
// 		assert.Equal(t, true, getB(t, Q("1")))
// 		assert.Equal(t, true, getB(t, Q("T")))
// 		assert.Equal(t, true, getB(t, Q("true")))
// 		assert.Equal(t, true, getB(t, Q("True")))
// 		assert.Equal(t, true, getB(t, Q("TRUE")))
// 		assert.Equal(t, false, getB(t, Q("0")))
// 		assert.Equal(t, false, getB(t, Q("F")))
// 		assert.Equal(t, false, getB(t, Q("false")))
// 		assert.Equal(t, false, getB(t, Q("False")))
// 		assert.Equal(t, false, getB(t, Q("FALSE")))
// 	}
// 	{
// 		q, _ := FromYaml(`foo:
//   bar: true`)
// 		assert.True(t, q.Any())

// 		// Exists
// 		assert.Equal(t, true, q.Yaml("foo.bar").O())
// 		assert.Equal(t, true, getB(t, q.Yaml("foo.bar")))

// 		// Doesn't exist
// 		assert.Equal(t, nil, q.Yaml("foo.foo").O())
// 		assert.Equal(t, false, getB(t, q.Yaml("foo.foo")))
// 	}
// }

// func TestI(t *testing.T) {
// 	assert.Equal(t, 2, getI(t, Q(2)))
// 	assert.Equal(t, 2, getI(t, Q("2")))
// 	assert.Equal(t, 3, Q([]int{2, 3}).At(1).O())
// 	assert.Equal(t, 3, getI(t, Q([]int{2, 3}).At(1)))
// 	assert.Equal(t, 3, getI(t, Q([]string{"2", "3"}).At(1)))
// 	assert.Equal(t, 1, getI(t, Q(true)))
// 	assert.Equal(t, 0, getI(t, Q(false)))
// }

// func TestM(t *testing.T) {
// 	{
// 		q, _ := FromYaml(`foo:
//   bar: true`)
// 		assert.True(t, q.Any())

// 		// Key exists
// 		assert.Equal(t, map[string]interface{}{"bar": true}, getM(t, q.Yaml("foo")))

// 		// Key doesn't exist
// 		assert.Equal(t, map[string]interface{}{}, getM(t, q.Yaml("bar")))
// 	}
// 	{
// 		// Convert with simple cast
// 		data := map[string]interface{}{"1": "one", "2": "two", "3": "three"}
// 		q := Q(data)
// 		assert.Equal(t, 3, q.Len())
// 		assert.Equal(t, data, getM(t, q))
// 	}
// 	{
// 		// Convert list of KeyVal into a map
// 		q := Q([]string{"foo=bar"})
// 		m := getM(t, q.Map(func(x O) O {
// 			return A(x.(string)).Split("=").YamlKeyVal()
// 		}))
// 		assert.Equal(t, map[string]interface{}{"foo": "bar"}, m)
// 	}
// }

// func TestInts(t *testing.T) {
// 	assert.Equal(t, []int{1, 2, 3}, getInts(t, Q([]int{1, 2, 3})))
// }

// func TestAAMap(t *testing.T) {
// 	{
// 		q, _ := FromYaml(`foo:
//   bar: frodo`)
// 		assert.True(t, q.Any())

// 		// Key exists
// 		assert.Equal(t, map[string]string{"bar": "frodo"}, getAAMap(t, q.Yaml("foo")))

// 		// Key doesn't exist
// 		assert.Equal(t, map[string]string{}, getAAMap(t, q.Yaml("bar")))
// 	}
// 	{
// 		// Get map of string to string
// 		q, _ := FromYaml(`foobar:
//   labels:
//     name: one
//     meta: two
//     data: three`)
// 		assert.True(t, q.Any())
// 		expected := map[string]string{"name": "one", "meta": "two", "data": "three"}
// 		assert.Equal(t, expected, getAAMap(t, q.Yaml("foobar.labels")))
// 	}
// 	{
// 		// Ints as keys
// 		q, _ := FromYaml(`foobar:
//   labels:
//     1: one
//     2: two
//     3: 3`)
// 		assert.True(t, q.Any())
// 		expected := map[string]string{"1": "one", "2": "two", "3": "3"}
// 		assert.Equal(t, expected, getAAMap(t, q.Yaml("foobar.labels")))
// 	}
// }

// func TestASAMap(t *testing.T) {
// 	{
// 		q := Q(map[string][]string{"bar": []string{"frodo"}})
// 		assert.Equal(t, map[string][]string{"bar": []string{"frodo"}}, getASAMap(t, q))
// 	}
// 	{
// 		q, err := FromYaml("foo:\n  bar:\n  - frodo")
// 		assert.Nil(t, err)
// 		assert.True(t, q.Any())
// 		assert.Equal(t, map[string][]string{"bar": []string{"frodo"}}, getASAMap(t, q.Yaml("foo")))
// 	}
// }

// func TestS(t *testing.T) {
// 	{
// 		q, _ := FromYaml(`items:
//   - name: one
//   - name: two
//   - name: three`)
// 		assert.True(t, q.Any())
// 		expected := []interface{}{
// 			map[string]interface{}{"name": "one"},
// 			map[string]interface{}{"name": "two"},
// 			map[string]interface{}{"name": "three"},
// 		}
// 		assert.Equal(t, expected, q.Yaml("items").S())
// 	}
// 	{
// 		// Not found
// 		q, _ := FromYaml(`items:
//   - name: one
//   - name: two
//   - name: three`)
// 		assert.True(t, q.Any())
// 		assert.Equal(t, N(), q.Yaml(""))
// 	}
// }

// func TestSAMap(t *testing.T) {
// 	// test map[interface{}]interface{}
// 	{
// 		q := Q([]interface{}{
// 			map[interface{}]interface{}{
// 				"name": "name1",
// 				"url":  "url1",
// 			},
// 			map[interface{}]interface{}{
// 				"name": "name2",
// 				"url":  "url2",
// 			},
// 		})
// 		expected := []map[string]interface{}{
// 			{
// 				"name": "name1",
// 				"url":  "url1",
// 			},
// 			{
// 				"name": "name2",
// 				"url":  "url2",
// 			},
// 		}
// 		assert.Equal(t, expected, getSAMap(t, q))
// 	}
// 	{
// 		q, _ := FromYaml(`items:
//   - name: one
//   - name: two
//   - name: three`)
// 		assert.True(t, q.Any())
// 		expected := []map[string]interface{}{
// 			{"name": "one"},
// 			{"name": "two"},
// 			{"name": "three"},
// 		}
// 		assert.Equal(t, expected, getSAMap(t, q.Yaml("items")))
// 	}
// 	{
// 		q, _ := FromYaml(`items:
//   - name: one
//   - name: two
//   - name: three`)
// 		assert.True(t, q.Any())
// 		expected := []map[string]interface{}{}
// 		assert.Equal(t, expected, getSAMap(t, q.Yaml("foo")))
// 	}
// }

// func TestSAAMap(t *testing.T) {
// 	// test map[interface{}]string
// 	{
// 		q := Q([]interface{}{
// 			map[interface{}]string{
// 				"name": "name1",
// 				"url":  "url1",
// 			},
// 			map[interface{}]string{
// 				"name": "name2",
// 				"url":  "url2",
// 			},
// 		})
// 		expected := []map[string]string{
// 			{
// 				"name": "name1",
// 				"url":  "url1",
// 			},
// 			{
// 				"name": "name2",
// 				"url":  "url2",
// 			},
// 		}
// 		assert.Equal(t, expected, getSAAMap(t, q))
// 	}
// 	// test map[interface{}]interface{}
// 	{
// 		q := Q([]interface{}{
// 			map[interface{}]interface{}{
// 				"name": "name1",
// 				"url":  "url1",
// 			},
// 			map[interface{}]interface{}{
// 				"name": "name2",
// 				"url":  "url2",
// 			},
// 		})
// 		expected := []map[string]string{
// 			{
// 				"name": "name1",
// 				"url":  "url1",
// 			},
// 			{
// 				"name": "name2",
// 				"url":  "url2",
// 			},
// 		}
// 		assert.Equal(t, expected, getSAAMap(t, q))
// 	}

// 	// slice of string to string map
// 	{
// 		q, _ := FromYaml(`items:
//   - name: one
//   - name: two
//   - name: three`)
// 		assert.True(t, q.Any())
// 		expected := []map[string]string{
// 			{"name": "one"},
// 			{"name": "two"},
// 			{"name": "three"},
// 		}
// 		assert.Equal(t, expected, getSAAMap(t, q.Yaml("items")))
// 	}

// 	// slice of string to int map
// 	{
// 		q, _ := FromYaml(`items:
//   - name: 1
//   - name: 2
//   - name: three`)
// 		assert.True(t, q.Any())
// 		expected := []map[string]string{
// 			{"name": "1"},
// 			{"name": "2"},
// 			{"name": "three"},
// 		}
// 		assert.Equal(t, expected, getSAAMap(t, q.Yaml("items")))
// 	}
// }

// func TestStrs(t *testing.T) {
// 	{
// 		// slice of string to int map
// 		q, _ := FromYaml(`items:
//   - name: 1
//   - name: 2
//   - name: three`)
// 		assert.True(t, q.Any())
// 		expected := []string{}
// 		assert.Equal(t, expected, q.Yaml("frodo.baggins").Strs())
// 	}
// 	{
// 		q := Q([]string{"one"})
// 		assert.Equal(t, "one", q.At(0).A())
// 		assert.Equal(t, []string{"one"}, q.Strs())
// 	}
// 	{
// 		assert.Equal(t, []string{"1", "2", "3"}, Q([]interface{}{"1", "2", "3"}).Strs())
// 	}
// }

// func getB(t *testing.T, q *Queryable) bool {
// 	result, err := q.B()
// 	assert.Nil(t, err)
// 	return result
// }

// func getI(t *testing.T, q *Queryable) int {
// 	result, err := q.I()
// 	assert.Nil(t, err)
// 	return result
// }

// func getInts(t *testing.T, q *Queryable) []int {
// 	result, err := q.Ints()
// 	assert.Nil(t, err)
// 	return result
// }

// func getM(t *testing.T, q *Queryable) map[string]interface{} {
// 	result, err := q.M()
// 	assert.Nil(t, err)
// 	return result
// }

// func getAAMap(t *testing.T, q *Queryable) map[string]string {
// 	result, err := q.AAMap()
// 	assert.Nil(t, err)
// 	return result
// }

// func getASAMap(t *testing.T, q *Queryable) map[string][]string {
// 	result, err := q.ASAMap()
// 	assert.Nil(t, err)
// 	return result
// }

// func getSAMap(t *testing.T, q *Queryable) []map[string]interface{} {
// 	result, err := q.SAMap()
// 	assert.Nil(t, err)
// 	return result
// }

// func getSAAMap(t *testing.T, q *Queryable) []map[string]string {
// 	result, err := q.SAAMap()
// 	assert.Nil(t, err)
// 	return result
// }
