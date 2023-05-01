# pointer [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/pointer.svg)](https://pkg.go.dev/github.com/carlmjohnson/pointer)

Some generic Go pointer helpers:

```go
var strptr *string
fmt.Println(pointer.Deref(strptr) == "") // prints true
fmt.Println(pointer.Coalesce(strptr, "hello, world")) // prints "hello, world"
newptr := pointer.From("meaning of life") // take a pointer to a string, wow!
strptr = pointer.First(strptr, newptr) // return the first non-nil pointer
```
