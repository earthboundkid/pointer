# pointer [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/pointer.svg)](https://pkg.go.dev/github.com/carlmjohnson/pointer)

Some generic Go pointer helpers:

```go
// Start with a null pointer
var strptr *string

// pointer.Deref safely dereferences a nil pointer into its empty value
fmt.Println(pointer.Deref(strptr) == "") // prints true

// pointer.Coalesce lets us specify a default value for a nil pointer
fmt.Println(pointer.Coalesce(strptr, "hello, world")) // prints "hello, world"

// We can create a pointer to a string or other primitive type with pointer.New
newptr := pointer.New("meaning of life") // takes a pointer to a string, wow!

// pointer.First returns the first pointer that isn't nil.
strptr = pointer.First(strptr, newptr) // returns newptr
```
