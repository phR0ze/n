package ntrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var one bool
var trace []*Entry
var skipfiles []string

func TestCallerTrace(t *testing.T) {

	// Test without any skips
	AppFunc1()
	assert.Len(t, trace, 10)

	// Test with single skip file
	skipfiles = append(skipfiles, "tracelog2_test")
	AppFunc1()
	assert.Len(t, trace, 8)

	// Test with multiple skip files
	skipfiles = append(skipfiles, "tracelog1_test")
	AppFunc1()
	assert.Len(t, trace, 6)
	assert.Equal(t, "ntrace.AppFunc3", trace[0].Func)
	assert.Equal(t, "ntrace.AppFunc2", trace[1].Func)
	assert.Equal(t, "ntrace.AppFunc1", trace[2].Func)
	assert.Equal(t, "ntrace.TestCallerTrace", trace[3].Func)
}

func TestCallerTraceOne(t *testing.T) {
	one = true
	skipfiles = []string{}

	// Test with single skip file
	AppFunc1()
	assert.Len(t, trace, 1)
	assert.Equal(t, "ntrace.LogFunc4", trace[0].Func)

	// Test with single skip file
	skipfiles = append(skipfiles, "tracelog2_test")
	AppFunc1()
	assert.Len(t, trace, 1)
	assert.Equal(t, "ntrace.LogFunc2", trace[0].Func)

	// // Test with multiple skip files
	skipfiles = append(skipfiles, "tracelog1_test")
	AppFunc1()
	assert.Len(t, trace, 1)
	assert.Equal(t, "ntrace.AppFunc3", trace[0].Func)
}

// Trace funcs
func AppFunc1() {
	AppFunc2()
}
func AppFunc2() {
	AppFunc3()
}
func AppFunc3() {
	LogFunc1()
}
