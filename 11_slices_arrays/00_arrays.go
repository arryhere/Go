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

	var v5 = v3[2:5]
	var v6 = v3[:]
	var v7 = v3[5:]
	var v8 = v3[:5]

	fmt.Println("v1:", v1) // v1: [0 0 0 0 0 0 0 0 0 0]
	fmt.Println("v2:", v2) // v2: [1 2 3 4 5 0 0 0 0 0]
	fmt.Println("v3:", v3) // v3: [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("v5:", v5) // v5: [3 4 5]
	fmt.Println("v6:", v6) // v6: [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("v7:", v7) // v7: [6 7 8 9 10]
	fmt.Println("v8:", v8) // v8: [1 2 3 4 5]

	fmt.Println("length v8:", len(v8))   // length v8: 5
	fmt.Println("capacity v8:", cap(v8)) // capacity v8: 10 [Explanation: v8 sliced from v3, where lenght of v3 is 10]
}
