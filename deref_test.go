package pointer_test

import (
	"fmt"

	"github.com/carlmjohnson/pointer"
)

func ExampleCoalesce() {
	var np *int
	fmt.Println(pointer.Coalesce(np, 1))
	np = new(int)
	fmt.Println(pointer.Coalesce(np, 1))
	// Output:
	// 1
	// 0
}

func ExampleDeref() {
	var np *int
	fmt.Println(pointer.Deref(np))
	one := 1
	np = &one
	fmt.Println(pointer.Deref(np))
	// Output:
	// 0
	// 1
}
