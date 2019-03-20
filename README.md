# n
***Numerable*** is a collection of missing Go convenience functions reminiscent of Ruby/C#. I love
the elegance of Ruby's plethera of descriptive chainable methods while C# has some awesome deferred
execution. Why not stay with Ruby or C# then? Well I like Go's ability to generate a single
statically linked binary, Go's concurrancy model, Go's performance, Go's package ecosystem and Go's
tool chain. This package was created to reduce the friction I had adopting Go as my primary
language of choice. ***n*** is intended to reduce the coding verbosity required by Go via
convenience functions and the Numerable types.

https://godoc.org/github.com/phR0ze/n

## Table of Contents
* [Numerable](#Numerable)
  * [Requirements](#requirements)
* [Background](#background)
  * [Performance](#performance)
    * [Go vs Python](#go-vs-python)
    * [Pure Reflection - 10 cost](#pure-reflection-10x-cost)
    * [Generic Slice - 18 cost](#generic-slice-18x-cost)
  * [Deferred Execution](#deferred-execution)
    * [Iterator Pattern](#iterator-pattern)

# Numerable <a name="numerable"></a>
***Numerable*** provide a way to generically handle various types in Go with the convenience
methods you would expect similar to Ruby or C#, making life a little easier. Since I'm using
Reflection to accomplish this it obviously comes at a cost, which in some cases isn't worth it.
However, as found in many cases, the actual work being done far out ways the bookkeeping overhead
incurred with the use of reflection. Other times the speed and convenience of not having to
re-implement a Delete or Contains function for the millionth time far out weighs the performance
cost.

## Numerable Requirements <a name="requirements"></a>
The Numerable interface and types implementing it are designed to accomplish the following
requirements:
* ***Chaining*** - the ability to call additional methods via a returned reference to the type
* ***Brevity*** - keep the naming as concise as possible while not infringing on clarity
* ***Clarity*** - keep the naming as unambiguous as possible while not infringing on brevity
* ***Performance*** - keep convenience functions as performant as possible while calling out significant costs
* ***Speed*** - provide convenience function parity with other rapid development languages
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

## Language Popularity <a name="language-popularity"></a>
* https://www.tiobe.com/tiobe-index/
* https://stackify.com/popular-programming-languages-2018/

## Performance <a name="performance"></a>
Performance is a concern in handling generics as the Golang inventors rightly pointed out. Go was
targeted as a systems language yet also noted as being a rapid development language. Certainly in my
experience it is being used in place of rapid development languages such as Ruby, Python and C# but
also Java as well. Generics are so vital to rapid development that a 10x cost may be worth it.

### Benchmarks Game <a name="benchmarks-game"></a>

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

### Reflection Assisted - 2x cost <a name="reflection-assisted-2x-cost"></a>
By storing items as a `reflect.ValueOf` we can use reflection to assist in handling generic
types but then type assert for all common types to provide those types with only a 2x cost vs
the 10x cost that falling back on pure reflection for unhandled types will cost.

6 nines slice costs 0.02ns while this implementation costs 0.02ns:
```golang
func (q *QSlice) Append(item interface{}) *QSlice {
	if q.Nil() {
		nq := Slicef(item)
		if !nq.Nil() {
			*q = *nq
		}
	} else {
		switch slice := q.v.Interface().(type) {
		case []int:
			if x, ok := item.(int); ok {
				slice = append(slice, x)
			} else {
				panic(fmt.Sprintf("can't insert type '%T' into []string", item))
			}
		default:
			panic("unsupported")
			item := reflect.ValueOf(item)
			*q.v = reflect.Append(*q.v, item)
		}
	}
	return q
}
```

### Pure Reflection - 10x cost <a name="pure-reflection-10x-cost"></a>
Storing the items as a `reflect.ValueOf` is the most elegant and obvious way of
doing this and as eveyone knows incurs the standard 10x reflection cost.

6 nines slice costs 0.01ns while this implementation costs 0.10ns:
```golang
func (q *QSlice) Append(items ...interface{}) *QSlice {
	if len(items) > 0 {
		if q.Nil() {
			*q = *(Slicef(items...))
		} else {
			for i := 0; i < len(items); i++ {
				item := reflect.ValueOf(items[i])
				*q.v = reflect.Append(*q.v, item)
			}
		}
	}
	return q
}
```

### Generic Slice - 18x cost <a name="generic-slice-18x-cost"></a>
Storing the items as a `[]interface{}` avoids the upfront 10x reflection cost but then requires
looping over the entire set of items and performing a type assertion on each to return the final
typed slice which resulted in an 18x cost even though reflection wasn't used. 

6 nines slice costs 0.01ns while this implementation costs 0.20ns:
```golang
func (q *QSlice) Append(items ...interface{}) *QSlice {
	if q.Nil() {
		*q = *(Slicef(items))
	} else {
		for _, item := range items {
			q.o = append(q.o, item)
		}
	}
	return q
}
```

## Deferred Execution <a name="deferred-execution"></a>
C# has some excellent defferred execution and the concept is really slick. I haven't found a great
need for it yet and thus haven't gotten around to it, but it's fun to research how it's done.

### Iterator Pattern <a name="iterator-pattern"></a>
Since Numerable is fundamentally based on the notion of iterables, iterating over collections, that
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
