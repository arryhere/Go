package variables

import (
	"fmt"
	"reflect"
)

func Variables_Datatypes() {

	var v1 bool = true
	var v2 string = "Hello"
	var v3 int = 55
	var v4 float32 = 55.55
	var v5 float64 = 55.55555555 // recommended over float32
	var v6 complex64 = complex(5, 5)
	var v7 complex128 = complex(5, 5) // recommended over complex64

	fmt.Println("v1: ", v1, ";", "typeOf: ", reflect.TypeOf(v1)) // v1:  true ; typeOf:  bool
	fmt.Println("v2: ", v2, ";", "typeOf: ", reflect.TypeOf(v2)) // v2:  Hello ; typeOf:  string
	fmt.Println("v3: ", v3, ";", "typeOf: ", reflect.TypeOf(v3)) // v3:  55 ; typeOf:  int
	fmt.Println("v4: ", v4, ";", "typeOf: ", reflect.TypeOf(v4)) // v4:  55.55 ; typeOf:  float32
	fmt.Println("v5: ", v5, ";", "typeOf: ", reflect.TypeOf(v5)) // v5:  55.55555555 ; typeOf:  float64
	fmt.Println("v6: ", v6, ";", "typeOf: ", reflect.TypeOf(v6)) // v6:  (5+5i) ; typeOf:  complex64
	fmt.Println("v7: ", v7, ";", "typeOf: ", reflect.TypeOf(v7)) // v7:  (5+5i) ; typeOf:  complex128

}
