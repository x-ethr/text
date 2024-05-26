package text_test

import (
	"fmt"
)

func ExampleLowercase() {
	v := "Field-Name"

	fmt.Println(str.Lowercase(v))
	// Output: field-name
}

func ExampleTitle() {
	// v represents a name that should be otherwise capitalized (titled)
	v := "jacob b. sanders"

	fmt.Println(str.Title(v))
	// Output: Jacob B. Sanders
}

func ExampleVariadic() {
	v := str.Dereference(nil, func(o str.Options) {
		o.Log = true
	})

	fmt.Println(v)
	// Output:

	pointer := str.Pointer("example")
	v = str.Dereference(pointer, func(o str.Options) {
		o.Log = true
	})

	fmt.Println(v)
	// Output: example
}

func ExampleDereference() {
	// initialize a string pointer of underlying value "example"
	pointer := str.Pointer("example")

	// establish variable "v" of type string
	v := str.Dereference(pointer, func(o str.Options) {
		o.Log = true // log if the pointer is nil
	})

	fmt.Println(v)
	// Output: example
}

func ExamplePointer() {
	// create a pointer of type string with reference value: "example"
	pointer := str.Pointer("example", func(o str.Options) {
		o.Log = true // log if the string value is an empty string
	})

	fmt.Println(*(pointer))
	// Output: example
}
