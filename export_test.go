package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	{
		q := Q("one")
		assert.Equal(t, "o", q.At(0).A())
	}
	{
		q := Q([]string{"one"})
		assert.Equal(t, "one", q.At(0).A())
		assert.Equal(t, []string{"one"}, q.Strs())
	}
	{
		assert.Equal(t, "1", Q("1").A())
	}
}

func TestI(t *testing.T) {
	assert.Equal(t, 1, Q(1).I())
}

func TestM(t *testing.T) {
	{
		data := map[string]interface{}{"1": "one", "2": "two", "3": "three"}
		q := Q(data)
		assert.Equal(t, 3, q.Len())
		assert.Equal(t, data, q.M())
	}
}

func TestInts(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Q([]int{1, 2, 3}).Ints())
}

func TestStrs(t *testing.T) {
	{
		q := Q([]string{"one"})
		assert.Equal(t, "one", q.At(0).A())
		assert.Equal(t, []string{"one"}, q.Strs())
	}
	{
		assert.Equal(t, []string{"1", "2", "3"}, Q([]interface{}{"1", "2", "3"}).Strs())
	}
}

func TestAAMap(t *testing.T) {
	{
		// Get map of string to string
		q := FromYAML(`foobar:
  labels:
    name: one
    meta: two
    data: three`)
		assert.True(t, q.Any())
		expected := map[string]string{"name": "one", "meta": "two", "data": "three"}
		assert.Equal(t, expected, q.YAML("foobar.labels").AAMap())
	}
	{
		// Ints as keys
		q := FromYAML(`foobar:
  labels:
    1: one
    2: two
    3: 3`)
		assert.True(t, q.Any())
		expected := map[string]string{"1": "one", "2": "two", "3": "3"}
		assert.Equal(t, expected, q.YAML("foobar.labels").AAMap())
	}
}

func TestS(t *testing.T) {
	{
		q := FromYAML(`items:
  - name: one
  - name: two
  - name: three`)
		assert.True(t, q.Any())
		expected := []interface{}{
			map[string]interface{}{"name": "one"},
			map[string]interface{}{"name": "two"},
			map[string]interface{}{"name": "three"},
		}
		assert.Equal(t, expected, q.YAML("items").S())
	}
}

func TestSAAMap(t *testing.T) {
	{
		q := FromYAML(`items:
  - name: one
  - name: two
  - name: three`)
		assert.True(t, q.Any())
		expected := []map[string]string{
			{"name": "one"},
			{"name": "two"},
			{"name": "three"},
		}
		assert.Equal(t, expected, q.YAML("items").SAAMap())
	}
}

func TestCastToTypeOf(t *testing.T) {
	//typof := []int{1, 2, 3}
	//obj := []interface{}{4}[0]
	//CastToTypeOf(typof, obj)
}
