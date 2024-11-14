package interfaces

import "fmt"

type IMathematics interface {
	SumOfTwoIntegers(first_value int, second_value int) (result int)
}

type Mathematics struct{}

func (m Mathematics) SumOfTwoIntegers(first_value int, second_value int) (result int) {
	return first_value + second_value
}

func Interface_Arguments() {
	v := Mathematics{}

	fmt.Println("SumOfTwoIntegers:", v.SumOfTwoIntegers(44, 55)) // SumOfTwoIntegers: 99
}
