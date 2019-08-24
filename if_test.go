package n

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func BenchmarkIf_Do(t *testing.B) {
	results := []int{}
	for _, i := range Range(0, 999999) {
		If(i%2 == 0).Do(func(x int) { results = append(results, x) }, i)
	}
	assert.Equal(t, 10, results[5])
}

func BenchmarkIf_DoGoStyle(t *testing.B) {
	results := []int{}
	for _, i := range Range(0, 999999) {
		if i%2 == 0 {
			results = append(results, i)
		}
	}
	assert.Equal(t, 10, results[5])
}

func TestIf(t *testing.T) {

	// Simple error checks
	{
		assert.Equal(t, true, If(true, errors.New("foo")).State)
		assert.Equal(t, "foo", If(true, errors.New("foo")).Error.Error())
	}

	// Simple state checks
	{
		assert.Equal(t, true, If(true).State)
		assert.Equal(t, nil, If(true).Error)
		assert.Equal(t, false, If(false).State)
		assert.Equal(t, nil, If(false).Error)
	}
}

func TestDo(t *testing.T) {

	// simple modulus
	{
		results := []int{}
		for _, i := range Range(0, 10) {
			If(i%2 == 0).Do(func(x int) { results = append(results, x) }, i)
		}
		assert.Equal(t, []int{0, 2, 4, 6, 8, 10}, results)
	}

	// one param
	{
		assert.Equal(t, "incorrect number of parameters for the given function", If(true).Do(func(x string, y bool) string { return x }, "foo").Error.Error())
		assert.Equal(t, "foo", If(true).Do(func(x string, y bool) string { return x }, "foo", true).Return[0])
		assert.Equal(t, "foo", If(true).Do(func(x string) string { return x }, "foo").Return[0])
	}

	// return string
	{
		assert.Equal(t, nil, If(true).Do(func() interface{} { return "foo" }).Error)
		assert.Equal(t, false, If(true).Do(func() string { return "foo" }).State)
		assert.Len(t, If(true).Do(func() string { return "foo" }).Return, 1)
		assert.Equal(t, "foo", If(true).Do(func() string { return "foo" }).Return[0])
	}

	// return bool
	{
		assert.Equal(t, true, If(true).Do(func() bool { return true }).State)
		assert.Equal(t, nil, If(true).Do(func() bool { return true }).Error)
		assert.Len(t, If(true).Do(func() bool { return true }).Return, 1)
		assert.Equal(t, true, If(true).Do(func() bool { return true }).Return[0])
	}

	// return error
	{
		assert.Equal(t, "err", If(true).Do(func() error { return errors.Errorf("err") }).Error.Error())
		assert.Len(t, If(true).Do(func() (string, error) { return "foo", errors.Errorf("err") }).Return, 2)
		assert.Equal(t, "err", If(true).Do(func() (string, error) { return "foo", errors.Errorf("err") }).Error.Error())
		assert.Equal(t, "foo", If(true).Do(func() (string, error) { return "foo", errors.Errorf("err") }).Return[0])
		assert.Equal(t, "err", If(true).Do(func() (string, error) { return "foo", errors.Errorf("err") }).Return[1].(error).Error())
	}
}

func TestReset(t *testing.T) {
	assert.Equal(t, false, If(true).Reset().State)
	assert.Equal(t, nil, If(true, errors.New("foo")).Reset().Error)
}
