package input

import "fmt"

func Input() {
	fmt.Print("Enter your first value: ")
	var first_value string
	fmt.Scanln(&first_value) // spaces or newline not included

	fmt.Print("Enter your second value: ")
	var second_value string
	fmt.Scanln(&second_value) // spaces or newline not included

	fmt.Println("Your first_value is:", first_value)
	fmt.Println("Your second_value is:", second_value)
}
