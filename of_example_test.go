package pointer_test

import (
	"fmt"

	"github.com/carlmjohnson/pointer"
)

func ExampleFrom() {
	strptr1 := pointer.From("meaning of life")
	strptr2 := pointer.From("meaning of life")
	fmt.Println(strptr1 != strptr2)
	fmt.Println(*strptr1 == *strptr2)

	intp1 := pointer.From(42)
	intp2 := pointer.From(42)
	fmt.Println(intp1 != intp2)
	fmt.Println(*intp1 == *intp2)

	type MyFloat float64
	fp := pointer.From[MyFloat](42)
	fmt.Println(fp != nil)
	fmt.Println(*fp == 42)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
