package slices_arrays

import "fmt"

func Make() {

	v1 := []int{}
	v2 := make([]int, 0)

	fmt.Println("v1:", v1) // v1: []
	fmt.Println("v2:", v2) // v2: []

	/*
		Prefer []int{} for readability unless you’re explicitly relying on make for additional clarity or consistency in your codebase.
	*/

	v3 := make([]int, 5)
	fmt.Println("v3:", v3) // v3: [0 0 0 0 0]

}
