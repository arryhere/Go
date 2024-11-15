package slices_arrays

import "fmt"

func Arrays() {
	var v1 = [10]int{}
	var v2 = [10]int{1, 2, 3, 4, 5}
	var v3 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	/*
		var v4 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

		index 10 is out of bounds (>= 10)
		index 11 is out of bounds (>= 10)
		index 12 is out of bounds (>= 10)
	*/

	fmt.Println("v1:", v1) // v1: [0 0 0 0 0 0 0 0 0 0]
	fmt.Println("v2:", v2) // v2: [1 2 3 4 5 0 0 0 0 0]
	fmt.Println("v3:", v3) // v3: [1 2 3 4 5 6 7 8 9 10]
}
