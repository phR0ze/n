# nub
Collection of missing Go helper functions reminiscent of C#'s IEnumerable methods

## Thoughts
https://golang.org/pkg/container/list/
https://golang.org/src/container/list/list.go
https://ewencp.org/blog/golang-iterators/index.html

### Iterators
```Golang
for x := range Foo.Iter() {
    // do something with x
}

func (n *Nub) ForEachInt(pred func(int)) {
    for _, val := range n.Data {
        pred(val)
    }
}

nub.ForEach()
```

### Collection Interface
Missing helper methods

```Golang
func From(interface{}) Nub {}
func FromSlice([]interface{}) Nub {}
func FromMap() Nub {}
func FromInt() Nub {}
func FromInt64() Nub {}
...
```
* ICollection
    * Add
    * Clear
    * Contains
    * Remove
* IEnumerable
    * All
    * Any
    * Append
    * Average
    * Contains
    * Count
    * Distinct
    * ElementAt
    * Except
    * First
    * GroupBy
    * Intersect
    * Join
    * Last
    * Max
    * Min
    * Prepend
    * Reverse
    * Select
    * SelectMany
    * Single
    * Skip
    * Sum
    * Take
    * TakeWhile
    * ToList
    * ToMap
    * Union
    * Where
    

* IList
    * Add
    * AddRange
    * Clear
    * Contains
    * Find
    * FindAll
    * ForEach
    * IndexOf
    * Insert
    * LastIndexOf
    * Remove
    * RemoveAny
    * RemoveAt
    * Reverse
    * Sort