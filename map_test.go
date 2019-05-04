package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewMap
//--------------------------------------------------------------------------------------------------
func ExampleNewMap() {
	fmt.Println(NewMap(map[string]interface{}{"k": "v"}))
	// Output: &map[k:v]
}

func TestNewMap(t *testing.T) {

	// string interface
	{
		m := map[string]interface{}{"k": "v"}
		assert.Equal(t, NewStringMap(m), NewMap(m))
	}

	// StringMap
	{
		m := NewStringMap()
		m.Set("k", "v")
		assert.Equal(t, m, NewMap(m))
	}
}
