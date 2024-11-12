package operators

import "fmt"

func Logical_Operators() {

	var v1 bool = true
	var v2 bool = false

	logical_and := v1 && v2
	logical_or := v1 || v2
	logical_not := !v1

	fmt.Printf("v1[%t] && v2[%t] = %t\n", v1, v2, logical_and) // v1[true] && v2[false] = false
	fmt.Printf("v1[%t] || v2[%t] = %t\n", v1, v2, logical_or)  // v1[true] || v2[false] = true
	fmt.Printf("!v1[%t] = %t\n", v1, logical_not)              // !v1[true] = false
}
