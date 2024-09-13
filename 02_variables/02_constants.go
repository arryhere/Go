package variables

import (
	"fmt"
	"reflect"
)

func Constants() {

	const v1 string = "Hello" // constants do not have shorthand decleration

	fmt.Println("v1: ", v1, ";", "typeOf: ", reflect.TypeOf(v1))
}
