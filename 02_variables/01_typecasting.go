package variables

import (
	"fmt"
	"reflect"
)

func Typecasting() {

	var v1 float64 = 55.86

	var v2 int = int(v1)

	var v3 float64 = float64(v2)

	fmt.Println("v1: ", v1, ";", "typeOf: ", reflect.TypeOf(v1)) // v1:  55.86 ; typeOf:  float64
	fmt.Println("v2: ", v2, ";", "typeOf: ", reflect.TypeOf(v2)) // v2:  55 ; typeOf:  int
	fmt.Println("v3: ", v3, ";", "typeOf: ", reflect.TypeOf(v3)) // v3:  55 ; typeOf:  float64

}
