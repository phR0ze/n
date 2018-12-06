package ncli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCliOpts(t *testing.T) {
	{
		data := []string{}
		expected := map[string]interface{}{}
		assert.Equal(t, expected, ParseCliOpts(data))
	}
	{
		data := []string{"foo=bar"}
		expected := map[string]interface{}{"foo": "bar"}
		assert.Equal(t, expected, ParseCliOpts(data))
	}
	{
		data := []string{"k1=v1,k2=v2"}
		expected := map[string]interface{}{"k1": "v1", "k2": "v2"}
		assert.Equal(t, expected, ParseCliOpts(data))
	}
	{
		data := []string{"k1=v1", "k2=v2"}
		expected := map[string]interface{}{"k1": "v1", "k2": "v2"}
		assert.Equal(t, expected, ParseCliOpts(data))
	}
	{
		data := []string{"k1=v1", "k2=v2", "k3=v3,k4=v4"}
		expected := map[string]interface{}{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4"}
		assert.Equal(t, expected, ParseCliOpts(data))
	}
}
