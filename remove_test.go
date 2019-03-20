package n

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestTakeFirst(t *testing.T) {
// 	{
// 		// empty
// 		x, ok := Nil().TakeFirst()
// 		assert.False(t, ok)
// 		assert.Nil(t, x)
// 	}
// 	{
// 		// empty string
// 		x, ok := Q("").TakeFirst()
// 		assert.False(t, ok)
// 		assert.Nil(t, x)
// 	}
// 	{
// 		// non-iterable
// 		x, ok := Q(1).TakeFirst()
// 		assert.False(t, ok)
// 		assert.Nil(t, x)
// 	}
// 	{
// 		// ints
// 		q := Q([]int{1, 2, 3})
// 		x, ok := q.TakeFirst()
// 		assert.True(t, ok)
// 		assert.NotNil(t, x)
// 		assert.Equal(t, 1, x.(int))
// 		assert.Equal(t, []int{2, 3}, q.O())
// 		//q.Append(4)
// 		//assert.Equal(t, []int{2, 3, 4}, q.O())
// 	}
// 	{
// 		// strings
// 		q := Nil().Append("1", "2", "3")
// 		x, ok := q.TakeFirst()
// 		assert.True(t, ok)
// 		assert.NotNil(t, x)
// 		assert.Equal(t, "1", x.(string))
// 		assert.Equal(t, []string{"2", "3"}, q.Strs())
// 	}
// 	{
// 		// maps
// 		q := Q(map[string]interface{}{"1": "one", "2": "two", "3": "three"})
// 		assert.True(t, q.Any())
// 		x, ok := q.TakeFirst()
// 		assert.False(t, ok)
// 		assert.Nil(t, x)
// 	}
// 	{
// 		// bobs
// 		q := Nil().Append(bob{o: "3"})
// 		assert.True(t, q.Any())
// 		assert.Equal(t, bob{o: "3"}, q.First().O())
// 		x, ok := q.TakeFirst()
// 		assert.True(t, ok)
// 		assert.NotNil(t, x)
// 		assert.Equal(t, bob{o: "3"}, x.(bob))
// 		assert.False(t, q.Any())
// 	}
// }

// func TestTakeFirstCnt(t *testing.T) {
// 	{
// 		q := Q([]int{1, 2, 3})
// 		assert.Equal(t, []int{1, 2, 3}, q.O())
// 		assert.Equal(t, 3, q.Len())
// 		items := q.TakeFirstCnt(4)
// 		assert.Equal(t, []int{1, 2, 3}, items)
// 		assert.Equal(t, 0, q.Len())
// 	}
// }
