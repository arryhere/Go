package interfaces

import (
	"fmt"
	"math"
)

// interfaces
type Area interface {
	Area() float64
}

type Perimeter interface {
	Perimeter() float64
}

type Geometry interface {
	Area
	Perimeter
}

// Rectangle
type Rectangle struct {
	length  float64
	breadth float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.length * rectangle.breadth
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}

// Circle
type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

// Functions
func GetArea(Area Area) float64 {
	return Area.Area()
}

func GetGeometry(geometry Geometry) (float64, float64) {
	return geometry.Area(), geometry.Perimeter()
}

func Interfaces() {

	my_rectangle := Rectangle{length: 5, breadth: 10}
	my_circle := Circle{radius: 5}

	fmt.Println("my_rectangle, Area:", my_rectangle.Area()) // my_rectangle, Area: 50
	fmt.Println("my_circle, Area:", my_circle.Area())       // my_circle, Area: 78.53981633974483

	fmt.Println("my_rectangle, GetArea:", GetArea(my_rectangle)) // my_rectangle, GetArea: 50
	fmt.Println("my_circle, GetArea:", GetArea(my_circle))       // my_circle, GetArea: 78.53981633974483

	react_area, rect_peri := GetGeometry(my_rectangle)
	fmt.Println("my_rectangle - geometry:", react_area, rect_peri) // my_rectangle - geometry: 50 30

	cir_area, cir_peri := GetGeometry(my_circle)
	fmt.Println("my_circle - geometry:", cir_area, cir_peri) // my_circle - geometry: 78.53981633974483 31.41592653589793

}
