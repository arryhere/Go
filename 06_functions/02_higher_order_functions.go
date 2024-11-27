package functions

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}

func mul(x float64, y float64) float64 {
	return x * y
}

func aggregate(a float64, b float64, c float64, d func(float64, float64) float64) float64 {
	return d(d(a, b), c)
}

func Higher_Order_Functions() {

	res1 := aggregate(12, 2, 3, add)
	res2 := aggregate(12, 2, 3, mul)

	fmt.Println("res1:", res1) // res1: 17
	fmt.Println("res2:", res2) // res2: 72
}
