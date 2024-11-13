package interfaces

import "fmt"

type IMathematics interface {
	sumOfTwoIntegers(first_value int, second_value int) (result int)
}

type Mathematics struct{}

func (m Mathematics) sumOfTwoIntegers(first_value int, second_value int) (result int) {
	return first_value + second_value
}

func Interface_Arguments() {
	v := Mathematics{}

	fmt.Println("sumOfTwoIntegers:", v.sumOfTwoIntegers(44, 55)) // sumOfTwoIntegers: 99
}
