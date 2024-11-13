package interfaces

import "fmt"

type Car struct {
	color string
	price int
}

func (car Car) horn() string {
	return "Vromm Vromm"
}

func dum(my_car Car) {
	fmt.Println("my_car:", my_car)
}

func Interfaces() {

	v := Car{
		color: "red",
		price: 44,
	}

	dum(v)

}
