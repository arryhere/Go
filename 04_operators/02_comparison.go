package operators

import "fmt"

func Comparison_Operators() {

	var v1 int = 15
	var v2 int = 5

	equal := v1 == v2
	not_equal := v1 != v2
	less := v1 < v2
	less_or_equal := v1 <= v2
	greater := v1 < v2
	greater_or_equal := v1 <= v2

	fmt.Printf("v1[%d] != v2[%d] : %t\n", v1, v2, not_equal)        // v1[15] != v2[5] : true
	fmt.Printf("v1[%d] == v2[%d] : %t\n", v1, v2, equal)            // v1[15] == v2[5] : false
	fmt.Printf("v1[%d] <  v2[%d] : %t\n", v1, v2, less)             // v1[15] <  v2[5] : false
	fmt.Printf("v1[%d] <= v2[%d] : %t\n", v1, v2, less_or_equal)    // v1[15] <= v2[5] : false
	fmt.Printf("v1[%d] >  v2[%d] : %t\n", v1, v2, greater)          // v1[15] >  v2[5] : false
	fmt.Printf("v1[%d] >= v2[%d] : %t\n", v1, v2, greater_or_equal) // v1[15] >= v2[5] : false
}
