package error_handling

import (
	"fmt"
	"strconv"
)

func Parse_Int(str string) int {
	var value, error = strconv.Atoi(str)

	if error != nil {
		return -1
	} else {
		return value
	}
}

func Error_Handling() {
	res_1 := Parse_Int("111")
	res_2 := Parse_Int("44G")

	fmt.Println("res_1:", res_1) // res_1: 111
	fmt.Println("res_2:", res_2) // res_2: -1
}
