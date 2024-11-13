package functions

import "fmt"

func sum(v1 float64, v2 float64) float64 {
	return v1 + v2
}

func Functions() {
	fmt.Println("sum:", sum(44.5, 232.56)) // sum: 277.06
}
