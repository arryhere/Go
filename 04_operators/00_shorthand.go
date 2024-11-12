package operators

import (
	"fmt"
	"reflect"
)

func Shorthand_Operators() {

	var v1 int = 55 // normal assignment
	v2 := 65        // shorthand assignment

	// use shorthand assigment over normal assignment; it is a standard go practice

	fmt.Println("v1: ", v1, ";", "typeOf: ", reflect.TypeOf(v1)) // v1:  55 ; typeOf:  int
	fmt.Println("v2: ", v2, ";", "typeOf: ", reflect.TypeOf(v2)) // v2:  65 ; typeOf:  int

}
