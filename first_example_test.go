package pointer_test

import (
	"fmt"

	"github.com/carlmjohnson/pointer"
)

func ExampleFirst() {
	type config struct{ string }
	userInput := func() *config {
		return nil
	}
	someConfig := pointer.First(
		userInput(),
		&config{"default config"},
	)
	fmt.Println(someConfig)
	// Output:
	// &{default config}
}
