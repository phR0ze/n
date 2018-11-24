# nub
Collection of missing Go helper functions reminiscent Ruby

## Implemented

### Slice Functions
| Function  | Description                              | IntSlice | StrSlice | StrMapSlice |
| --------- | ---------------------------------------- | -------- | -------- | ----------- |
| NewTYPE   | Creates a new nub encapsulating the TYPE | 1        | 1        | 1           |
| Any       | Check if the slice has anything in it    | 1        | 1        | 1           |
| AnyWhere  | Match slice items against given lambda   | O        | O        | O           |
| Append    | Add items to the end of the slice        | 1        | 1        | 1           |
| At        | Get item using neg/pos index notation    | 1        | 1        | 1           |
| Clear     | Clear out the underlying slice           | 1        | 1        | 1           |
| Contains  |             | X        | X        |             |
| Count     |             |          |          |             |
| Distinct  |             | X        | X        |             |
| Each      |             |          |          |             |
| Find      |             |          |          |             |
| FindAny   |             |          |          |             |
| First     |             |          |          |             |
| Index     |             |          |          |             |
| Insert    |             |          |          |             |
| Join      |             |          |          |             |
| Last      |             |          |          |             |
| Prepend   |             |          |          |             |
| Del       |             |          |          |             |
| DelWhere  |             |          |          |             |
| Reverse   |             |          |          |             |
| Select    |             |          |          |             |
| Sort      |             |          |          |             |
| Shift     |             |          |          |             |
| ShiftCnt  |             |          |          |             |
| Take      |             |          |          |             |

## Thoughts
https://golang.org/pkg/container/list/
https://golang.org/src/container/list/list.go
https://ewencp.org/blog/golang-iterators/index.html