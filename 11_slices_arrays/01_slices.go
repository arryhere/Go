package slices_arrays

import "fmt"

func Slices() {
	a1 := []int{} // slice

	for i := 1; i <= 10; i++ {
		a1 = append(a1, i)
	}

	fmt.Println("a1:", a1) // a1: [1 2 3 4 5 6 7 8 9 10]

	a2 := [3]int{1, 2, 3} // array

	for i := 1; i <= 10; i++ {
		a1 = append(a2[:], i) // a2[:] Converting a2 array to a slice, as append works on slice only
	}

	fmt.Println("a1:", a1) // a1: [1 2 3 10]
}
