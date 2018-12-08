# nos
The `Nub OS` package is intended to provide system related helper functions for Golang commonly
found with other languages.

## Table of Contents
* [Functions](#functions)
  * [Copy(src, dst string) error](#copy)

## Functions <a name="functions"></a>

### Copy(src, dst string) error <a name="copy"></a>
```Go
// Copy copies src to dst recursively.
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist, but it's parent directory does.
func Copy(src, dst string) error
```

```Go
// CopyFile copies a single file from src to dst.
// The dst will be copied to if it is an existing directory.
// The dst will be a clone of the src if it doesn't exist, but it's parent directory does.
func CopyFile(src, dst string) (err error) {
```
