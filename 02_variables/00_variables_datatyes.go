package main

import (
	"fmt"
	"reflect"
)

func main() {

	var v1 bool = true
	var v2 string = "Hello"
	var v3 int = 55
	var v4 float32 = 55.55
	var v5 float64 = 55.55555555
	var v6 complex64 = complex(5, 5)
	var v7 complex128 = complex(5, 5)

	fmt.Println("v1: ", v1, ";", "typeOf: ", reflect.TypeOf(v1))
	fmt.Println("v2: ", v2, ";", "typeOf: ", reflect.TypeOf(v2))
	fmt.Println("v3: ", v3, ";", "typeOf: ", reflect.TypeOf(v3))
	fmt.Println("v4: ", v4, ";", "typeOf: ", reflect.TypeOf(v4))
	fmt.Println("v5: ", v5, ";", "typeOf: ", reflect.TypeOf(v5))
	fmt.Println("v6: ", v6, ";", "typeOf: ", reflect.TypeOf(v6))
	fmt.Println("v7: ", v7, ";", "typeOf: ", reflect.TypeOf(v7))
}
