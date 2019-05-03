package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Set
//--------------------------------------------------------------------------------------------------
func ExampleStringMap_Set() {
	fmt.Println(NewStringMap().Set("key", "value"))
	// Output: true
}

func TestStringMap_Set(t *testing.T) {
	// bool
	{
		m := NewStringMap()
		assert.Equal(t, true, m.Set("test", false))
		assert.Equal(t, false, m.Set("test", true))
		assert.Equal(t, true, (*m)["test"].(bool))
	}
}
