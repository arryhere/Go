package conditionals

import (
	"fmt"
)

func If_Else() {
	value := 4

	if value > 50 {
		fmt.Println("Value is > 50, value:", value)
	} else if value == 50 {
		fmt.Println("Value is == 50, value:", value)
	} else if value < 50 {
		fmt.Println("Value is < 50, value:", value) // Value is < 50, value: 4
	} else {
		fmt.Println("Value is wierd !")
	}

	if v := true; v {
		fmt.Println("v is true !")
	}

}
