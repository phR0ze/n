# nub
Collection of missing Go helper functions reminiscent Ruby

## Implemented
Ruby and C# both have excellent helper methods for collections which Go either lacks entirely
or has tucked away in various packages that are difficult for newbies to find and extermely
verbose to use.  I find it extremely tedious to continually re-implement simple basic functions
which is why I'm creating nub objects with helper functions. I've chosen a handful to implement
and skipped over others that can be accomplished in a simpler way with those i did choose.

### Slice Functions
| Function     | Description                                     | IntSlice | StrSlice | StrMapSlice |
| ------------ | ----------------------------------------------- | -------- | -------- | ----------- |
| NewTYPE      | Creates a new nub encapsulating the TYPE        | 1        | 1        | 1           |
| Any          | Check if the slice has anything in it           | 1        | 1        | 1           |
| AnyWhere     | Match slice items against given lambda          | 0        | 0        | 0           |
| Append       | Add items to the end of the slice               | 1        | 1        | 1           |
| At           | Get item using neg/pos index notation           | 1        | 1        | 1           |
| Clear        | Clear out the underlying slice                  | 1        | 1        | 1           |
| Contains     | Check if the slice contains the given item      | 1        | 1        | 1           |
| ContainsAny  | Check if the slice contains any given items     | 1        | 1        | 1           |
| Count        | Count items that match lambda result            | 0        | 0        | 0           |
| Del          | Delete item using neg/pos index notation        | 1        | 1        | 1           |
| DelWhere     | Delete the items that match the given lambda    | 0        | 0        | 0           |
| Each         | Execute given lambda for each item in slice     | 0        | 0        | 0           |
| Index        | Get the index of the item matchin the given     | 0        | 0        | 0           |
| Insert       | Insert an item into the underlying slice        | 0        | 0        | 0           |
| Join         | Join slice items as string with given delimiter | 1        | 1        | 0           |
| Len          | Get the length of the slice                     | 1        | 1        | 1           |
| M            | Materializes object invoking deferred execution | 1        | 1        | 1           |
| Prepend      | Add items to the begining of the slice          | 1        | 1        | 1           |
| Reverse      | Reverse the items                               | 0        | 0        | 0           |
| Sort         | Sort the items                                  | 1        | 1        | 0           |
| TakeFirst    | Remove and return the first item from the slice | 1        | 1        | 1           |
| TakeFirstCnt | Remove and return the first cnt items           | 1        | 1        | 1           |
| TakeLast     | Remove and return the last item from the slice  | 1        | 1        | 1           |
| TakeLastCnt  | Remove and return the last cnt items            | 1        | 1        | 1           |
| Uniq         | Ensure only uniq items exist in the slice       | 1        | 1        | 0           |
| Where        | Select the items that match the given lambda    | 0        | 0        | 0           |

## Thoughts
https://golang.org/pkg/container/list/
https://golang.org/src/container/list/list.go
https://ewencp.org/blog/golang-iterators/index.html