package functions

import "fmt"

func sum(v1 float64, v2 float64) float64 {
	return v1 + v2
}

func multiple_return() (string, string) {
	return "Hello", "World"
}

func Functions() {
	fmt.Println("sum:", sum(44.5, 232.56)) // sum: 277.06

	var t1, t2 = multiple_return()
	fmt.Println("multiple_return:", t1, t2)
}
