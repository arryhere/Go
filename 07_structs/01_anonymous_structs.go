package structs

import "fmt"

func Anonymous_Structs() {

	var v = struct {
		value_1 int
		value_2 string
		value_3 struct {
			value_4 int
			value_5 string
		}
	}{
		value_1: 55,
		value_2: "Hello",
		value_3: struct {
			value_4 int
			value_5 string
		}{
			value_4: 66,
			value_5: "World",
		},
	}

	fmt.Println("v:", v) // v: {55 Hello {66 World}}

}
