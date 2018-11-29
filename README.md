# nub
Collection of missing Go helper functions reminiscent of Ruby/C#. I love the
elegance of Ruby's short named plethera of chainable methods while C# has
some awesome deferred execution. I'm attempting to marry the two :)

## Table of Contents
* [References](#references)
  * [Iterator Pattern](#iterator-pattern)
* [Implemented](#implemented)

## References <a name="references"></a>

### Deferred Execution <a name="deferred-execution"></a>
TBD

### Iterator Pattern <a name="iterator-pattern"></a>
I implemented the iterator pattern based off of the iterator closure pattern disscussed
by blog https://ewencp.org/blog/golang-iterators/index.htm mainly for syntactic style;
other [sources](https://stackoverflow.com/questions/14000534/what-is-most-idiomatic-way-to-create-an-iterator-in-go)
indicates that the closure style iterator is about 3x slower than normal. However my own benchmarking was much closer:

Changing the order in which my benchmarks run seems to affect the time (caching?)  
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

## Implemented <a name="implemented"></a>
Ruby and C# both have excellent helper methods for collections which Go either lacks entirely
or has tucked away in various packages that are difficult for newbies to find and extermely
verbose to use.  I find it extremely tedious to continually re-implement simple basic functions
which is why I'm creating nub objects with helper functions. I've chosen a handful to implement
and skipped over others that I may come back to.

### Functions <a name="functions"></a>
Some functions only apply to particular underlying collection types as called out in the table.

**Key: '1' = Implemented, '0' = Not Implemented, 'blank' = NA**

| Function     | Description                                     | Slice | Map | Str |  Cust |
| ------------ | ----------------------------------------------- | ----- | ----| --- | ----- |
| A            | Creates queryable encapsulating empty string    |       |     | 1   |       |
| M            | Creates queryable encapsulating empty map       |       | 1   |     |       |
| S            | Creates queryable encapsulating empty map       | 1     |     |     |       |
| N            | Creates queryable encapsulating empty slice     | 1     |     |     |       |
| Q            | Creates queryable encapsulating the given TYPE  | 1     | 1   | 1   | 1     |
| Any          | Check if the queryable is not nil and not empty | 1     | 1   | 1   | 1     |
| AnyWhere     | Check if any match the given lambda             | 1     | 1   | 1   | 1     |
| Append       | Add items to the end of the collection          | 1     | 0   | 1   | 0     |
| At           | Return item at the given neg/pos index notation | 1     |     | 1   | 0     |
| Clear        | Clear out the underlying collection             | 1     | 1   | 1   | 0     |
| Contains     | Check that all given items are found            | 1     | 1   | 1   | 1     |
| ContainsAny  | Check that any given items are found            | 1     | 1   | 1   | 1     |
| Each         | Iterate over the queryable and execute actions  | 1     | 1   | 1   | 1     |
| Join         | Join slice items as string with given delimiter | 1     |     |     |       |
| Len          | Get the length of the collection                | 1     | 1   | 1   | 1     |
| Load         | Load YAML/JSON from file into queryable         |       | 1   |     |       |
| Map          | Manipulate the queryable data into a new form   | 1     | 1   | 1   | 1     |
| Merge        | Merge other queryables in priority order        | 0     | 0   | 0   | 0     |
| Set          | Set the queryable's encapsulated object         | 1     | 1   | 1   | 1     |
| TakeFirst    | Remove and return the first item                | 1     |     | 1   | 1     |
| TakeFirstCnt | Remove and return the first cnt items           | 0     | 0   | 0   | 0     |
| TakeLast     | Remove and return the last item                 | 0     | 0   | 0   | 0     |
| TakeLastCnt  | Remove and return the last cnt items            | 0     | 0   | 0   | 0     |
| TypeIter     | Is queryable iterable                           | 1     | 1   | 1   | 1     |
| TypeMap      | Is queryable reflect.Map                        | 1     | 1   | 1   | 1     |
| TypeStr      | Is queryable encapsualting a string             | 1     | 1   | 1   | 1     |
| TypeSlice    | Is queryable reflect.Array or reflect.Map       | 1     | 1   | 1   | 1     |
| TypeSingle   | Is queryable encapsualting a non-collection     | 1     | 1   | 1   | 1     |

### Materialization <a name="materialization"></a>
Materialization or processing deferred execution and converting to a usable type

| Function     | Description                                     | Return Type               |
| ------------ | ----------------------------------------------- | ------------------------- |
| A            | Materialize queryable into a string             | `string`                  |
| I            | Materialize queryable into an int               | `int`                     |
| M            | Materialize queryable to map                    | `map[string]interface{}`  |
| O            | Materialize queryable to interface{}            | `interface{}`             |
| Ints         | Materialize queryable into an int slice         | `[]int`                   |
| Strs         | Materialize queryable into a string slice       | `[]string`                |

### String Functions <a name="string-functions"></a>
| Function     | Description                                     | Done  |
| ------------ | ----------------------------------------------- | ----- |
| Split        | Split the string into a slice on delimiter      |       |



### Slice Functions
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
| M            | Materializes object invoking deferred execution | 0     | 1        | 1        | 1           |
| Prepend      | Add items to the begining of the slice          | 0     | 1        | 1        | 1           |
| Reverse      | Reverse the items                               | 0     | 0        | 0        | 0           |
| Sort         | Sort the items                                  | 0     | 1        | 1        | 0           |
| Uniq         | Ensure only uniq items exist in the slice       | 0     | 1        | 1        | 0           |
| Where        | Select the items that match the given lambda    | 0     | 0        | 0        | 0           |

### Map Functions
| Function     | Description                                     | IntMap | StrMap | ? |
| ------------ | ----------------------------------------------- | -------- | -------- | ----------- |
| NewTYPE      | Creates a new nub encapsulating the TYPE        | 0        | 1        | 0           |
| Load         | Load YAML/JSON from file                        | 0        | 1        | 0           |
| Add          | Add a new item to the underlying map            | 0        | 1        | 0           |
| Any          | Check if the map has anything in it             | 0        | 1        | 0           |
| Equals       | Check if the given map is equal to this map     | 0        | 1        | 0           |
| Len          | Get the length of the map                       | 0        | 1        | 0           |
| M            | Materializes object invoking deferred execution | 0        | 1        | 0           |
| Merge        | Merge other maps in, in priority order          | 0        | 1        | 0           |
| MergeNub     | Merge other nub maps in, in priority order      | 0        | 1        | 0           |
| Slice        | Get the slice indicated by the multi-key        | 0        | 1        | 0           |
| Str          | Get the str indicated by the multi-key          | 0        | 1        | 0           |
| StrMap       | Get the str map indicated by the multi-key      | 0        | 1        | 0           |
| StrMapByName | Get the str map by name and multi-key           | 0        | 1        | 0           |
| StrMapSlice  | Get the str map slice by the multi-key          | 0        | 1        | 0           |
| StrSlice     | Get the str slice by the multi-key              | 0        | 1        | 0           |

## Thoughts
https://golang.org/pkg/container/list/
https://golang.org/src/container/list/list.go
https://ewencp.org/blog/golang-iterators/index.html