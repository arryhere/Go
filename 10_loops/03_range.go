package loops

import "fmt"

func Range() {
	var ls = []float64{}

	ls = append(ls, 1.2)
	ls = append(ls, 2.3)
	ls = append(ls, 3.4)
	ls = append(ls, 4.5)
	ls = append(ls, 5.6)

	fmt.Println("ls:", ls)

	for i, e := range ls {
		fmt.Printf("i: [%v], e: [%v]\n", i, e)
	}
}
