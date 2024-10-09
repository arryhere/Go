package input

import "fmt"

func Input() {
	fmt.Print("Enter your value: ")

	var value string
	fmt.Scanln(&value) // spaces or newline not included

	fmt.Println("Your value is:", value)
}
