package error_handling

import (
	"fmt"
	"strconv"
)

func ParseInt(str string) int {
	var value, error = strconv.Atoi(str)

	if error != nil {
		return -1
	} else {
		return value
	}
}

func Error_Handling() {
	res_1 := ParseInt("111")
	res_2 := ParseInt("44G")

	fmt.Println("res_1:", res_1) // res_1: 111
	fmt.Println("res_2:", res_2) // res_2: -1
}
