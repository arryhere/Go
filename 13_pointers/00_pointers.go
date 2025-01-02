package pointers

import "fmt"

func Pointers() {
	var p *int
	fmt.Println("p:", p) // p: <nil>

	var q int = 10
	p = &q
	fmt.Println("p:", p)   // p: 0x1400000e110
	fmt.Println("*p:", *p) // *p: 10

	*p = 100
	fmt.Println("q:", q)   // q: 100
	fmt.Println("p:", p)   // p: 0x1400000e110
	fmt.Println("*p:", *p) // *p: 100
}
