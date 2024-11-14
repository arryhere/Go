package functions

import "fmt"

func Sum(v1 float64, v2 float64) float64 {
	return v1 + v2
}

func MultipleReturn() (string, string) {
	return "Hello", "World"
}

func Functions() {
	fmt.Println("Sum:", Sum(44.5, 232.56)) // Sum: 277.06

	var t1, t2 = MultipleReturn()
	fmt.Println("MultipleReturn:", t1, t2) // MultipleReturn: Hello World
}
