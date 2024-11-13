package interfaces

import (
	"fmt"
	"math"
)

// interfaces
type Area interface {
	area() float64
}

type Perimeter interface {
	perimeter() float64
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

func (rectangle Rectangle) area() float64 {
	return rectangle.length * rectangle.breadth
}

func (rectangle Rectangle) perimeter() float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}

// Circle
type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

// Functions
func getArea(area Area) float64 {
	return area.area()
}

func getGeometry(geometry Geometry) (float64, float64) {

	return geometry.area(), geometry.perimeter()

}

func Interfaces() {

	my_rectangle := Rectangle{length: 5, breadth: 10}
	my_circle := Circle{radius: 5}

	fmt.Println("my_rectangle, area:", my_rectangle.area()) // my_rectangle, area: 50
	fmt.Println("my_circle, area:", my_circle.area())       // my_circle, area: 78.53981633974483

	fmt.Println("my_rectangle, getArea:", getArea(my_rectangle)) // my_rectangle, getArea: 50
	fmt.Println("my_circle, getArea:", getArea(my_circle))       // my_circle, getArea: 78.53981633974483

	react_area, rect_peri := getGeometry(my_rectangle)
	fmt.Println("my_rectangle - geometry:", react_area, rect_peri) // my_rectangle - geometry: 50 30

	cir_area, cir_peri := getGeometry(my_circle)
	fmt.Println("my_circle - geometry:", cir_area, cir_peri) // my_circle - geometry: 78.53981633974483 31.41592653589793

}
