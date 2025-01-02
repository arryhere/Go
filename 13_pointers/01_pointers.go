package pointers

import "fmt"

func mutate(input *string, value string) {
	*input = value
}

func Pointers_1() {
	x := "Hello"
	y := &x

	fmt.Println("x:", x) // x: Hello

	mutate(y, "Hello World")

	fmt.Println("x:", x) // x: Hello World
}
