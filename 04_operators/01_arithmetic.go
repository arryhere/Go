package operators

import "fmt"

func Arithmetic_Operators() {

	var v1 int = 15
	var v2 int = 5

	add := v1 + v2
	sub := v1 - v2
	mul := v1 * v2
	div := v1 / v2
	mod := v1 % v2

	fmt.Printf("v1[%d] + v2[%d] = %d\n", v1, v2, add)   // v1[15] + v2[5] = 20
	fmt.Printf("v1[%d] - v2[%d] = %d\n", v1, v2, sub)   // v1[15] - v2[5] = 10
	fmt.Printf("v1[%d] * v2[%d] = %d\n", v1, v2, mul)   // v1[15] * v2[5] = 75
	fmt.Printf("v1[%d] / v2[%d] = %d\n", v1, v2, div)   // v1[15] / v2[5] = 3
	fmt.Printf("v1[%d] mod v2[%d] = %d\n", v1, v2, mod) // v1[15] mod v2[5] = 0
}
