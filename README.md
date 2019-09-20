# n
***Nub*** is a collection of missing Go convenience functions reminiscent of Ruby/C#. I love
the elegance of Ruby's plethera of descriptive chainable methods while C# has some awesome deferred
execution. Why not stay with Ruby or C# then? Well I like Go's ability to generate a single
statically linked binary, Go's concurrancy model, Go's performance, Go's package ecosystem and Go's
tool chain. This package was created to reduce the friction I had adopting Go as my primary
language of choice. ***n*** is intended to reduce the coding verbosity required by Go via
convenience functions and the Nub types.

**Requires Go 1.13 and only supports POSIX systems**

https://godoc.org/github.com/phR0ze/n

## Table of Contents
* [Nub](#Nub)
  * [Requirements](#requirements)
* [Background](#background)
  * [Resources](#resources)
  * [Language Popularity](#language-popularity)
  * [Language Benchmarks](#language-benchmarks)
  * [Generic Performance](#generic-performance)
    * [Custom Native - 0x cost](#custom-native-0x-cost)
    * [Pure Reflection - 9x cost](#pure-reflection-9x-cost)
    * [Slice of interface{} - 14 cost](#slice-of-interface-14x-cost)
    * [Reflection Assisted - 6.83x cost](#reflection-assisted-6.83x-cost)
  * [Deferred Execution](#deferred-execution)
    * [Iterator Pattern](#iterator-pattern)

# Nub <a name="nub"></a>
***Nub*** provides a way to generically handle various types in Go with the convenience
methods you would expect similar to Ruby or C#, making life a little easier. I've gone to great
lengths to implement Nub types for all common Go types and only fall back on Reflection
for custom types. This means that in many cases no Reflection is used at all. In the cases where
Reflection is used it obviously comes at a cost, which in some cases won't be worth it. However
even then as found in many cases, the actual work being done far out ways the bookkeeping overhead
incurred with the use of reflection. Other times the speed and convenience of not having to
re-implement a Delete or Contains function for the millionth time far out weighs the performance
cost.

## Requirements <a name="requirements"></a>
The Nub types have been designed to accomplish the following requirements:

* ***Chaining*** - the ability to call additional methods via a returned reference to the type
* ***Brevity*** - keep the naming as concise as possible while not infringing on clarity
* ***Clarity*** - keep the naming as unambiguous as possible while not infringing on brevity
* ***Performance*** - keep convenience functions as performant as possible while calling out significant costs
* ***Speed*** - provide convenience function parity with rapid development languages
* ***Comfort*** - use naming and concepts in similar ways to popular languages

# Background <a name="background"></a>
Efficiency is paramount when developing. We want to develop quickly, be able to pick up someone
else's code and understand it quickly and yet still have our code execute quickly. The industry uses
terms like ***Code Readability***, ***Code Reuse***, ***Code Maintainablity***, ***Code Clarity***,
***Code Performance*** to describe this. These best practices have developed over a long history
swinging wildly from one end of the spectrum to the other and back again. 

Development started in the 1950's with super low level langauges and little readability, clarity or
maintainability but awesome performance (relatively speaking). This was far left on the efficiency
spectrum of performance vs rapid development. Coding was tedious and laborious. Then came the systems
level languages like C (1970's) and C++ (1980's) that began shifting a little more to the right with
more abstraction and convenience functions to reuse algorithums and code thus higher development
speed. This worked so well that higher level languages with even more abstractions and even more
convenience functions were created like Java (1990's), Ruby (1990's), Python (1990's), C# (2000's)
etc... all chasing this dream of optimal efficiency but swinging away from performance toward rapid
development on the far right. The inventors of Golang felt that the trade off with current languages
was unacceptable and designed Go to address the problem. To their credit they were able to get a
pretty good mix between performance and speed of development.

## Resources <a name="resources"></a>
* Go is not a Good language - http://yager.io/programming/go.html
* http://speakmy.name/2014/09/14/modifying-interfaced-go-struct/

## Language Popularity <a name="language-popularity"></a>
* https://www.tiobe.com/tiobe-index/

## Language Benchmarks <a name="language-benchmarks"></a>
* https://www.techempower.com/benchmarks/
* https://www.quora.com/Why-use-Rails-if-NET-Core-is-so-much-faster-in-benchmarks
* https://www.edureka.co/blog/golang-vs-python/
* https://getstream.io/blog/switched-python-go/
* https://benchmarksgame-team.pages.debian.net/benchmarksgame/faster/go.html
* https://benchmarksgame-team.pages.debian.net/benchmarksgame/faster/python3-go.html
* https://benchmarksgame-team.pages.debian.net/benchmarksgame/faster/ruby.html
* https://codeburst.io/javascript-vs-ruby-vs-python-which-is-the-best-language-for-your-startup-e072b14bebc7
* https://benchmarksgame-team.pages.debian.net/benchmarksgame/which-programs-are-fast.html
* http://ece.uprm.edu/~nayda/Courses/Icom5047F06/Papers/paper4.pdf

| Test            | Py 3.7.2 | Go 1.12 | C# 2.2.1 | Ruby 2.6 | Go vs Py |
| --------------- | -------- | ------- | -------- | -------- | -------- |
| Binary Trees    | 81.74s   | 26.94s  | 7.73s    | 64.97s   | 3.03x    | 
| Fannkuch-redux  | 482.90s  | 18.07s  | 17.46s   | 545.63s  | 26.72x   | 
| Fasta           | 63.18s   | 2.07s   | 2.27s    | 63.32s   | 30.52x   | 
| K-nucleotide    | 73.60s   | 11.98s  | 5.48s    | 189.81s  | 6.14x    | 
| Mandlebrot      | 265.56s  | 5.50s   | 5.54s    | 452.81s  | 48.28x   |
| N-Body          | 819.56s  | 21.19s  | 21.41s   | 387.73s  | 38.68x   | 
| Pidigits        | 3.33s    | 2.03s   | 0.84s    | 1.71s    | 1.64x    | 
| Regex-Redux     | 17.28s   | 29.13s  | 2.22s    | 23.58    | -1.67x   | 
| Reverse Comple  | 16.19s   | 3.98s   | 2.99s    | 23.09s   | 4.05x    | 
| Spectral-Norm   | 170.52s  | 3.94s   | 4.07s    | 237.96s  | 43.28x   | 

## Generic Performance <a name="generic-performance"></a>
Performance is a concern in handling generics as the Golang inventors rightly pointed out. Go was
targeted as a systems language yet also noted as being a rapid development language. Certainly in my
experience it is being used in place of rapid development languages such as Ruby, Python and C# but
also Java as well. Generics are so vital to rapid development that a 10x cost may be worth it when
that is required. In the following sections I examine different implementation workarounds for this
whole in the language and the cost associated with those implementations.

To do this I'll be implementing a `Map` function that will for testing purposes also include the
creation of the initial slice from a set of seed data as well as iterating over the seed data
using a user given lambda to manipulate the data and return a property of the original object then
to return this new slice of data as a native type.

I'll use these helper types and functions to generate the test data:
```golang
type Integer struct { Value int}
func Range(min, max int) []Integer {
	result := make([]Integer, max-min+1)
	for i := range result {
		result[i] = Bob{min + i
	}
	return result
}
```

### Custom Native - 0x cost <a name="custom-native-0x-cost"></a>
To set a base of comparision we'll implement the desired functionality assuming we know the
types i.e. there is nothing that can be reused in this case and a developer must implement
their own basic functionality for the new type.

Results from 3 runs after cache warm up:
```
BenchmarkSlice_CustomNative-16    	2000000000	         0.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_CustomNative-16    	2000000000	         0.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_CustomNative-16    	2000000000	         0.01 ns/op	       0 B/op	       0 allocs/op
```

```golang
func BenchmarkSlice_CustomNative(t *testing.B) {
	seedData := RangeInteger(0, 999999)

	// Select the actual values out of the custom object
	lambda := func(x Integer) int {
		return x.Value
	}

	// Because we are assuming the developer isn't reusing anything we know the types
	// and can use them directly
	ints := []int{}
	for i := range seedData {
		ints = append(ints, lambda(seedData[i]))
	}

	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}
```

### Pure Reflection - 9x cost <a name="pure-reflection-9x-cost"></a>
The first and most obvious way to deal with the short coming is to work around it using reflection.
Typically whenever reflection comes into the picture there is a 10x cost associated with it. 
However I wanted to run some of my own benchmarks as a fun exercise.

***9x*** hit from 3 runs after cache warm up:
```
BenchmarkSlice_PureReflection-16    	2000000000	         0.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_PureReflection-16    	2000000000	         0.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_PureReflection-16    	2000000000	         0.09 ns/op	       0 B/op	       0 allocs/op
```

```golang
func BenchmarkSlice_PureReflection(t *testing.B) {
	seedData := RangeInteger(0, 999999)

	// Select the actual values out of the custom object
	lambda := func(x interface{}) interface{} {
		return x.(Integer).Value
	}

	// Use reflection to determine the Kind this would model creating a reference
	// to a new type e.g. NewRefSlice(seed). Because we used reflection we don't
	// need to convert the seed data we can just use the given slice object the
	// way it is.
	var v2 reflect.Value
	v1 := reflect.ValueOf(seedData)
	k1 := v1.Kind()

	// We do need to validate that is a slice type before working with it
	if k1 == reflect.Slice {
		for i := 0; i < v1.Len(); i++ {
			result := lambda(v1.Index(i).Interface())

			// We need to create a new slice based on the result type to store all the results
			if !v2.IsValid() {
				typ := reflect.SliceOf(reflect.TypeOf(result))
				v2 = reflect.MakeSlice(typ, 0, v1.Len())
			}

			// Appending the native type to a reflect.Value of slice is more complicated as we
			// now have to convert the result type into a reflect.Value before appending
			// type asssert the native type.
			v2 = reflect.Append(v2, reflect.ValueOf(result))
		}
	}

	// Now convert the results into a native type
	// Because we created the native slice type as v2 we can simply get its interface and cast
	ints := v2.Interface().([]int)
	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}
```

### Slice of interface{} - 14x cost <a name="slice-of-interface-14x-cost"></a>
Now we'll try using the native `[]interface{}` type to accomplish the same task of appending
to a given slice a set of values from another slice. I was initially shocked over the results of this
as I fully expected this to be faster until I looked a little closer. Because we are working with a
slice of `interface{}` we have go to thank in that we can't simply cast any slice type to a slice of
interface. Instead we have to iterate over it and pass in each item to the new `[]interface{}` type.
This means that we need to iterate over the entire slice once to convert to a slice of interface then
again to execute the lambda over each item then again to convert it into something usable. That is
3 complete loops and we still had to use reflectdion to get into an initial known state.

***14x*** hit from 3 runs after cache warm up:
```
BenchmarkSlice_SliceOfInterface-16    	2000000000	         0.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_SliceOfInterface-16    	2000000000	         0.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_SliceOfInterface-16    	2000000000	         0.14 ns/op	       0 B/op	       0 allocs/op
```

```golang
func BenchmarkSlice_SliceOfInterface(t *testing.B) {
	seedData := RangeInteger(0, 999999)

	// Select the actual values out of the custom object
	lambda := func(x interface{}) interface{} {
		return x.(Integer).Value
	}

	// Use reflection to determine the Kind this would model creating a reference
	// to a new type e.g. NewSlice(seed). Once we've done that we need to convert
	// it into a []interface{}
	v1 := reflect.ValueOf(seedData)
	k1 := v1.Kind()

	// We do need to validate that is a slice type before working with it
	g1 := []interface{}{}
	if k1 == reflect.Slice {
		for i := 0; i < v1.Len(); i++ {
			g1 = append(g1, v1.Index(i).Interface())
		}
	}

	// Now iterate and execute the lambda and create the new result
	results := []interface{}{}
	for i := range g1 {
		results = append(results, lambda(g1[i]))
	}

	// Now convert the results into a native type
	ints := []int{}
	for i := range results {
		ints = append(ints, results[i].(int))
	}
	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}
```

### Reflection Assisted - 6.83x cost <a name="reflection-assisted-6.83x-cost"></a>
Reflection assisted is the notion that we can can develop native types to support
common types e.g. `[]int` and provide helper methods for them and fall back on 
reflection for custom types that are not yet implemented.

***8x*** hit from 3 runs after cache warm up for a custom type using reflection:
```
BenchmarkSlice_RefSlice-16    	2000000000	         0.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_RefSlice-16    	2000000000	         0.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_RefSlice-16    	2000000000	         0.07 ns/op	       0 B/op	       0 allocs/op
```

```golang
func BenchmarkSlice_RefSlice(t *testing.B) {
	ints := NewSlice(RangeInteger(0, 999999)).Map(func(x O) O {
		return x.(Integer).Value
	}).ToInts()
	assert.Equal(t, 2, ints[2])
	assert.Equal(t, 99999, ints[99999])
}
```

***6x*** hit from 3 runs after cache warm up using IntSlice Nub type:
```
BenchmarkSlice_IntSlice-16    	2000000000	         0.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_IntSlice-16    	2000000000	         0.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice_IntSlice-16    	2000000000	         0.06 ns/op	       0 B/op	       0 allocs/op
```

```golang
func BenchmarkSlice_IntSlice(t *testing.B) {
	ints := NewSlice(Range(0, 999999)).Map(func(x O) O {
		return x.(int) + 1
	}).ToInts()
	assert.Equal(t, 3, ints[2])
	assert.Equal(t, 100000, ints[99999])
}
```

## Deferred Execution <a name="deferred-execution"></a>
C# has some excellent defferred execution and the concept is really slick. I haven't found a great
need for it yet and thus haven't gotten around to it, but it's fun to research how it's done.

### Iterator Pattern <a name="iterator-pattern"></a>
Since Nub is fundamentally based on the notion of iterables, iterating over collections, that
was the first challenge to solve. How do you generically iterate over all primitive Go types.

I implemented the iterator pattern based off of the iterator closure pattern disscussed by blog
https://ewencp.org/blog/golang-iterators/index.htm mainly for syntactic style.  Some
[sources](https://stackoverflow.com/questions/14000534/what-is-most-idiomatic-way-to-create-an-iterator-in-go)
indicate that the closure style iterator is about 3x slower than normal. However my own benchmarking
was much closer coming in at only 33% hit. Even at 3x slower I'm willing to take that hit for
convenience in most case.

Changing the order in which my benchmarks run seems to affect the time (caching?).
At any rate on average I'm seeing only about a 33.33% performance hit. 33% in nested large
data sets may impact performance in some cases but I'm betting in most cases performance
will be dominated by actual work and not looping overhead.

```bash
# 36% slower to use Each function
BenchmarkEach-16               	       1	1732989848 ns/op
BenchmarkArrayIterator-16      	       1	1111445479 ns/op
BenchmarkClosureIterator-16    	       1	1565197326 ns/op

# 25% slower to use Each function
BenchmarkArrayIterator-16      	       1	1210185264 ns/op
BenchmarkClosureIterator-16    	       1	1662226578 ns/op
BenchmarkEach-16               	       1	1622667779 ns/op

# 30% slower to use Each function
BenchmarkClosureIterator-16    	       1	1695826796 ns/op
BenchmarkArrayIterator-16      	       1	1178876072 ns/op
BenchmarkEach-16               	       1	1686159938 ns/op
```
