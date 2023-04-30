# pointer [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/pointer.svg)](https://pkg.go.dev/github.com/carlmjohnson/pointer)

Some generic Go pointer helpers:

```go
var strptr *string
fmt.Println(pointer.Deref(strptr) == "") // prints true
fmt.Println(pointer.Coalesce(strptr, "hello, world")) // prints "hello, world"
strptr = pointer.From("meaning of life")
```
