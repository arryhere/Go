package structs

import "fmt"

type Quadrilateral struct {
	length  float64
	breadth float64
}

func (quadrilateral Quadrilateral) area() float64 {
	return quadrilateral.length * quadrilateral.breadth
}

func Struct_Methods() {
	var q = Quadrilateral{
		length:  5,
		breadth: 8,
	}

	fmt.Println("area:", q.area()) // area: 40
}
