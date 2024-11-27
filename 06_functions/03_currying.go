package functions

import "fmt"

func _mul(x float64, y float64) float64 {
	return x * y
}

func _add(x float64, y float64) float64 {
	return x + y
}

func selfOperation(fx func(x float64, y float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return fx(x, x)
	}
}

/* */

func run(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x + y + z
		}
	}
}

func Currying() {
	fx_square := selfOperation(_mul)
	res_fx_square := fx_square(5)

	fx_double := selfOperation(_add)
	res_fx_double := fx_double(5)

	fmt.Println("res_fx_square:", res_fx_square) // res_fx_square: 25
	fmt.Println("res_fx_double:", res_fx_double) // res_fx_double: 10

	/* */

	res := run(5)(5)(7)
	fmt.Println("res:", res) // res: 17
}
