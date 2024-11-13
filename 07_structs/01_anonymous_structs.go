package structs

import "fmt"

func Anonymous_Structs() {

	var v = struct {
		value_1 int
		value_2 string
	}{
		value_1: 55,
		value_2: "Hello",
	}

	fmt.Println("v:", v) // v: {55 Hello}

}
