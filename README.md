# n
***n*** is a collection of missing Go convenience functions reminiscent of Ruby/C#. I love the
elegance of Ruby's plethera of descriptive chainable methods while C# has some awesome deferred
execution. Why not stay with Ruby or C# then? Well I like Go's ability to generate a single
statically linked binary, Go's concurrancy model, Go's performance, Go's package ecosystem and Go's
tool chain. This package was created to reduce the friction I had adopting Go as my primary
language of choice. ***n*** is intended to reduce the coding verbosity required by Go via
convenience functions and the Queryable types.

## Table of Contents
* [Queryable](#queryable)
  * [Types](#types)
    * [Constructors](#constructors)
  * [Functions](#functions)
  * [QSlice](#qslice)

  * [Methods](#queryable-methods)
  * [Exports](#queryable-exports)
* [QStr](#qstr)
  * [QStr Functions](#qstr-functions)
  * [QStr Exports](#qstr-exports)
  * [QStr Methods](#qstr-methods)
* [Background](#background)
  * [Array Type](#array-type)
  * [Deferred Execution](#deferred-execution)
    * [Iterator Pattern](#iterator-pattern)

# Queryable <a name="queryable"></a>
***Queryable*** provide a way to generically handle collections in Go with the convenience
methods you would expect similar to Ruby or C# making life a little easier. Since I'm using
Reflection to accomplish this it obviously comes at a cost, which in some cases isn't worth it.
However, as found in many cases, the actual work being done far out ways the bookkeeping overhead
incurred with the use of reflection. Other times the speed and convenience of not having to
re-implement a Delete or Contains function for the millionth time far out weighs the performance
cost.

Queryable Requirements:
* ***Chaining*** - the ability to call additional array methods via a returned reference to the collection
* ***Brevity*** - keep the naming as concise as possible while not infringing on clarity
* ***Clarity*** - keep the naming as unambiguous as possible while not infringing on brevity
* ***Performance*** - keep convenience functions as performant as possible while calling out significant costs
* ***Development Speed*** - provide convenience function parity with other rapid development languages
* ***Comfort*** - use naming and concepts in similar ways to popular languages

## Types <a name="types"></a>
***n*** provides a number of types to assist in working with collections. 

| Type         | Description                                                                        |
| ------------ | ---------------------------------------------------------------------------------- |
| QSlice       | Provides a generic way to work with slice types providing convenience methods      |


| O            | O is an alias for interface{} to reduce verbosity                           |
| Queryable    | Chainable execution and is the heart of algorithm abstration layer          |
| Iterator     | Closure interator pattern implementation                                    |
| KeyVal       | Simple key value pair structure for iterating over map types                |

### Constructors <a name="constructors"></a>
Each collection type implementing the Queryable interface should have three ways to create it.

Example:
```golang
# The default constructor
$ slice := &QSlice{}

# The wrapper constructor
$ slice := Slice([]string{"test"})

# The variadic constructor
$ slice := Slicef("test1", "test2")
```

| Function  | Description                                                                   | Bench |
| --------- | ----------------------------------------------------------------------------- | ----- |
| Slice     | Instantiates a new QSlice optionally seeding it with the given obj if a Slice |       |
| Slicef    | Instantiates a new QSlice optionally seeding it with the given variadic obj   |       |


## Functions <a name="functions"></a>
***n*** provides a number of functions to assist in working with collection types.

| Function  | Description                                                                   | Bench |
| --------- | ----------------------------------------------------------------------------- | ----- |


| N            | Creates queryable encapsulating nil             | 1     |     |     |      |
| Q            | Creates queryable encapsulating the given TYPE  | 1     | 1   | 1   | 1    |


## QSlice <a name="qslice"></a>
Every language has the ability to collect contiguous items together in an ordered way. Typically
this is referred to as an array, list, or an arraylist. Languages typically include functions for
manipulating these ordered items such as adding, removing, sorting, counting etc... ***QList***,
which wraps a Go Slice, is intended to provide these missing conveniences.

Other language method references:
* C# - https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1.add?view=netframework-4.7.2
* Ruby - https://ruby-doc.org/core-2.6.0.preview2/Array.html
* Java - https://docs.oracle.com/javase/8/docs/api/java/util/List.html

| Method    | Description                                                                   | Bench |
| --------- | ----------------------------------------------------------------------------- | ----- |
| Append    | Append items to the end of the QSlice and return QSlice for chaining          | 9x    |

| Go     | slice = append(slice, item)
| C#     | list.Add(item)
| Python | list.append(item)
| Ruby   | list.append(item)
| Ruby   | list.push(item)

| append(item)           | Add a single element to the list
| extend(other)          | Add elements of the given list to the list
| insert(index, element) | Inserts element to the list
| remove(element)        | Removes Element from the List
| index()                | returns smallest index of element in list
| count()                | returns occurrences of element in a list
| pop()                  | Removes Element at Given Index
| reverse()              | Reverses a List
| sort()                 | sorts elements of a list
| copy()                 | Returns Shallow Copy of a List
| clear()                | Removes all Items from the List
| any()                  | Checks if any Element of an Iterable is True
| all()                  | returns true when all elements in iterable is true
| ascii()                | Returns String Containing Printable Representation
| bool()                 | Converts a Value to Boolean
| enumerate()            | Returns an Enumerate Object
| filter()               | constructs iterator from elements which are true
| iter()                 | returns iterator for an object
| list()                 | Function	creates list in Python
| len()                  | Returns Length of an Object
| max()                  | returns largest element
| min()                  | returns smallest element
| map()                  | Applies Function and Returns a List
| reversed()             | returns reversed iterator of a sequence
| slice()                | creates a slice object specified by range
| sorted()               | returns sorted list from a given iterable
| sum()                  | Add items of an Iterable
| zip()                  | Returns an Iterator of Tuples


## Methods <a name="methods"></a>
Some methods only apply to particular underlying collection types as called out in the table.

**Key: '1' = Implemented, '0' = Not Implemented, 'blank' = Unsupported, Bench nx = slowness**

| Function     | Description                                     | Slice | Map | Str | Cust | Bench |
| ------------ | ----------------------------------------------- | ----- | ----| --- | ---- | ----- |
| Any          | Check if the queryable is not nil and not empty | 1     | 1   | 1   | 1    | 1x    |
| AnyWhere     | Check if any match the given lambda             | 1     | 1   | 1   | 1    | 3x    |
| Append       | Add items to the end of the collection          | 1     |     | 1   | 1    | 10x   |
| At           | Return item at the given neg/pos index notation | 1     |     | 1   | 1    | 1x    |
| Clear        | Clear out the underlying collection             | 1     | 1   | 1   | 1    | 1x    |
| Contains     | Check that all given items are found            | 1     | 1   | 1   | 1    |       |
| ContainsAny  | Check that any given items are found            | 1     | 1   | 1   | 1    |       |
| Copy         | Copy the given obj into this queryable          | 1     | 1   | 1   | 1    | 1x    |
| Delete       | Delete all items that match the given obj       |       | 1   |     |      |       |
| DeleteAt     | Deletes the item at the given index location    | 1     | 1   | 1   | 1    | 1.10x |
| Each         | Iterate over the queryable and execute actions  | 1     | 1   | 1   | 1    | 1.33x |
| Join         | Join slice items as string with given delimiter | 1     |     |     |      |       |
| Len          | Get the length of the collection                | 1     | 1   | 1   | 1    |       |
| Load         | Load Yaml/JSON from file into queryable         |       | 1   |     |      |       |
| Map          | Manipulate the queryable data into a new form   | 1     | 1   | 1   | 1    |       |
| Merge        | Merge other queryables in priority order        | 0     | 0   | 0   | 0    |       |
| Set          | Set the queryable's encapsulated object         | 1     | 1   | 1   | 1    |       |
| TakeFirst    | Remove and return the first item                | 1     |     | 1   | 1    |       |
| TakeFirstCnt | Remove and return the first cnt items           | 0     | 0   | 0   | 0    |       |
| TakeLast     | Remove and return the last item                 | 0     | 0   | 0   | 0    |       |
| TakeLastCnt  | Remove and return the last cnt items            | 0     | 0   | 0   | 0    |       |
| TypeIter     | Is queryable iterable                           | 1     | 1   | 1   | 1    |       |
| TypeMap      | Is queryable reflect.Map                        | 1     | 1   | 1   | 1    |       |
| TypeStr      | Is queryable encapsualting a string             | 1     | 1   | 1   | 1    |       |
| TypeSlice    | Is queryable reflect.Array or reflect.Map       | 1     | 1   | 1   | 1    |       |
| TypeSingle   | Is queryable encapsualting a non-collection     | 1     | 1   | 1   | 1    |       |

## Exports <a name="exports"></a>
Exports process deferred execution and convert the result to a usable external type

| Function     | Description                                           | Return Type                |
| ------------ | ----------------------------------------------------- | -------------------------- |
| A            | Export queryable as a string                          | `string`                   |
| B            | Export queryable as a bool                            | `bool`                     |
| I            | Export queryable as an int                            | `int`                      |
| M            | Export queryable as a string map of interface{}       | `map[string]interface{}`   |
| O            | Export queryable as underlying type interface{}       | `interface{}`              |
| S            | Export queryable as a slice of interface{}            | `[]interface{}`            |
| Ints         | Export queryable as a slice of int                    | `[]int`                    |
| Strs         | Export queryable as a slice of string                 | `[]string`                 |
| AAMap        | Export queryable as a string map of string            | `map[string]string`        |
| ASAMap       | Export queryable as a string map of []string          | `map[string][]string`      |
| SAMap        | Export queryable as a slice of string map of...       | `[]map[string]interface{}` |
| SAAMap       | Export queryable as a slice of string map of...       | `[]map[string]string`      |

# QStr <a name="qstr"></a>
QStr implementes the Queryable Interface and integrates with other queryable types.  It provides a
plethora of convenience methods to work with string types.

## QStr Functions <a name="qstr-functions"></a>
| Function     | Description                                                                | Bench |
| ------------ | -------------------------------------------------------------------------- | ----- |
| A            | Instantiate a new QStr optionally seeding it with the given value          |       |

## QStr Exports <a name="qstr-exports"></a>
Exports process deferred execution and convert the result to a usable external type

| Function     | Description                                           | Return Type                |
| ------------ | ----------------------------------------------------- | -------------------------- |
| A            | Export QStr as a string                               | `string`                   |
| B            | Export QStr as bytes                                  | `[]byte`                   |
| Q            | Export QStr as a Queryable                            | `Queryable`                |

## QStr Methods <a name="qstr-methods"></a>
| Method       | Description                                                                | Bench |
| ------------ | -------------------------------------------------------------------------- | ----- |
| At           | Return run at the given neg/pos index notation                             | 1x    |
| Type         | Return the QType identifying this queryable type                           | 1x    |
| Contains     | Check that the given item is contained in the QStr                         |       |
| ContainsAll  | Check that all the given items are contained in the QStr                   |       |
| ContainsAny  | Check that any of the given items are contained in the QStr                |       |
| Empty        | Returns true if the pointer is nil, string is empty or whitespace only     |       |
| HasAnyPrefix | Checks if the string has any of the given prefixes                         |       |
| HasAnySuffix | Checks if the string has any of the given suffixes                         |       |
| HasPrefix    | Checks if the string has the given prefix                                  |       |
| Len          | Returns the length of the string                                           |       |
| Replace      | Wraps strings.Replace and allows for chaining and defaults                 |       |
| SpaceLeft    | Returns leading whitespace                                                 |       |
| SpaceRight   | Returns trailing whitespace                                                |       |
| Split        | Creates a new QSlice from the split string                                 |       |


## Slice Functions
| Function     | Description                                     | Slice | IntSlice | StrSlice | StrMapSlice |
| ------------ | ----------------------------------------------- | ----- | -------- | -------- | ----------- |
| NewTYPE      | Creates a new nub encapsulating the TYPE        | 1     | 1        | 1        | 1           |
| Any          | Check if the slice has anything in it           | 1     | 1        | 1        | 1           |
| AnyWhere     | Match slice items against given lambda          | 0     | 0        | 0        | 0           |
| Append       | Add items to the end of the slice               | 1     | 1        | 1        | 1           |
| At           | Get item using neg/pos index notation           | 0     | 1        | 1        | 1           |
| Clear        | Clear out the underlying slice                  | 0     | 1        | 1        | 1           |
| Contains     | Check if the slice contains the given item      | 0     | 1        | 1        | 1           |
| ContainsAny  | Check if the slice contains any given items     | 0     | 1        | 1        | 1           |
| Count        | Count items that match lambda result            | 0     | 0        | 0        | 0           |
| Del          | Delete item using neg/pos index notation        | 0     | 1        | 1        | 1           |
| DelWhere     | Delete the items that match the given lambda    | 0     | 0        | 0        | 0           |
| Each         | Execute given lambda for each item in slice     | 0     | 0        | 0        | 0           |
| Equals       | Check if the given slice is equal to this slice | 0     | 1        | 1        | 1           |
| Index        | Get the index of the item matchin the given     | 0     | 0        | 0        | 0           |
| Insert       | Insert an item into the underlying slice        | 0     | 0        | 0        | 0           |
| Join         | Join slice items as string with given delimiter | 0     | 1        | 1        | 0           |
| Len          | Get the length of the slice                     | 0     | 1        | 1        | 1           |
| E            | Exports object invoking deferred execution      | 0     | 1        | 1        | 1           |
| Prepend      | Add items to the begining of the slice          | 0     | 1        | 1        | 1           |
| Reverse      | Reverse the items                               | 0     | 0        | 0        | 0           |
| Sort         | Sort the items                                  | 0     | 1        | 1        | 0           |
| Uniq         | Ensure only uniq items exist in the slice       | 0     | 1        | 1        | 0           |
| Where        | Select the items that match the given lambda    | 0     | 0        | 0        | 0           |

## Map Functions
| Function     | Description                                     | IntMap | StrMap | ? |
| ------------ | ----------------------------------------------- | -------- | -------- | ----------- |
| NewTYPE      | Creates a new nub encapsulating the TYPE        | 0        | 1        | 0           |
| Load         | Load Yaml/JSON from file                        | 0        | 1        | 0           |
| Add          | Add a new item to the underlying map            | 0        | 1        | 0           |
| Any          | Check if the map has anything in it             | 0        | 1        | 0           |
| Equals       | Check if the given map is equal to this map     | 0        | 1        | 0           |
| Len          | Get the length of the map                       | 0        | 1        | 0           |
| M            | Exports object invoking deferred execution | 0        | 1        | 0           |
| Merge        | Merge other maps in, in priority order          | 0        | 1        | 0           |
| MergeNub     | Merge other nub maps in, in priority order      | 0        | 1        | 0           |
| Slice        | Get the slice indicated by the multi-key        | 0        | 1        | 0           |
| Str          | Get the str indicated by the multi-key          | 0        | 1        | 0           |
| StrMap       | Get the str map indicated by the multi-key      | 0        | 1        | 0           |
| StrMapByName | Get the str map by name and multi-key           | 0        | 1        | 0           |
| StrMapSlice  | Get the str map slice by the multi-key          | 0        | 1        | 0           |
| StrSlice     | Get the str slice by the multi-key              | 0        | 1        | 0           |

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

## Deferred Execution <a name="deferred-execution"></a>
C# has some excellent defferred execution and the concept is really slick. I haven't found a great
need for it yet and thus haven't gotten around to it, but it's fund to research how it's done.

### Iterator Pattern <a name="iterator-pattern"></a>
Since Queryable is fundamentally based on the notion of iterables, iterating over collections, that
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
