package functions

import "fmt"

func concat() func(string) string {
	s := ""
	return func(word string) string {
		if s == "" {
			s = s + word
		} else {
			s = s + " " + word
		}
		return s
	}
}

func CLosure() {
	fx := concat()

	fmt.Println(fx("Hello")) // Hello
	fmt.Println(fx("World")) // Hello World
	fmt.Println(fx("and"))   // Hello World and
	fmt.Println(fx("Hello")) // Hello World and Hello
	fmt.Println(fx("GO !"))  // Hello World and Hello GO !
}
