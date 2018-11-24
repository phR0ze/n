# nub
Collection of missing Go helper functions reminiscent Ruby

## Implemented

### Slice Functions
| Function  | Description                              | IntSlice | StrSlice | StrMapSlice |
| --------- | ---------------------------------------- | -------- | -------- | ----------- |
| NewTYPE   | Creates a new nub encapsulating the TYPE | 1        | 1        | 1           |
| Any       | Check if the slice has anything in it    | O        | O        | O           |
| AnyWhere  | Match slice items against given lambda   | O        | O        | O           |
| Append    | Add items to the end of the slice        | 1        | 1        | 1           |
| At        |             |          |          |             |
| Clear     |             |          |          |             |
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