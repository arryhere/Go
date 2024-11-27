package functions

import (
	"fmt"
	"strconv"
)

func message() {
	fmt.Println("[Defer]: Just a regular message!")
}

func Defer() {
	defer message() // will run only before the function returns

	fmt.Print("Enter your integer value: ")
	var int_value_input string
	fmt.Scanln(&int_value_input)

	int_value, int_value_input_error := strconv.ParseInt(int_value_input, 10, 0)

	if int_value_input_error != nil {
		fmt.Println("[Error]: int_value_input_error")
		return
	}

	if int_value > 0 {
		fmt.Printf("int_value: [%v],  value greater than 0\n", int_value)
		return
	} else if int_value < 0 {
		fmt.Printf("int_value: [%v], value less than 0\n", int_value)
		return
	} else {
		fmt.Printf("int_value: [%v], value is 0\n", int_value)
		return
	}
}
