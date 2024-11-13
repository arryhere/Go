package structs

import "fmt"

type BigContainer struct {
	value_1 int
	value_2 string
	value_3 bool
	value_4 SmallContainer
}

type SmallContainer struct {
	value_1 int
	value_2 string
}

func Structs() {
	var v BigContainer = BigContainer{
		value_1: 44,
		value_2: "Hello",
		value_3: true,
		value_4: SmallContainer{
			value_1: 55,
			value_2: "World",
		},
	}

	fmt.Println("v:", v) // v: {44 Hello true {55 World}}
}
