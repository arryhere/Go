package structs

import "fmt"

type Container_1 struct {
	value_1 int
	value_2 string
	value_3 bool
	Container_2
}

type Container_2 struct {
	value_4 int
	value_5 string
}

func Embedded_Structs() {

	v := Container_1{
		value_1: 55,
		value_2: "Hello",
		value_3: true,
		Container_2: Container_2{
			value_4: 66,
			value_5: "World",
		},
	}

	fmt.Println("v:", v) // v: {55 Hello true {66 World}}

	fmt.Println(v.value_1) // 55
	fmt.Println(v.value_4) // 66
}
