package main

import (
	"fmt"
	"math"
)

// interface
type Geometry interface {
	area() float64
	perim() float64
}

// struct
type React struct {
	width, height float64
}

type Circle struct {
	radius float64
}

//method calculate the area and perim for react and struct

func (r React) area() float64 {
	return r.width * r.height
}
func (r React) perim() float64 {
	result := 2*r.width + 2*r.height

	return math.Round(result)
}

func (c Circle) area() float64 {

	fmt.Println("float64 value of area ==> roundedToEven", math.Pi*c.radius*c.radius)
	return math.RoundToEven(math.Pi * c.radius * c.radius)
}

func (c Circle) perim() float64 {
	result := 2 * (math.Pi * c.radius)
	fmt.Println("float64 value of area ==> rounded", result)

	return math.Round(result)
}

//a function behalf as inheritor of the geometry interface

func measure(geometry Geometry) {
	fmt.Println(geometry)
	fmt.Println(geometry.area())
	fmt.Println(geometry.perim())
}

func main() {

	rectangle := React{height: 10, width: 2}
	measure(rectangle)

	circle := Circle{radius: 4}
	measure(circle)

}
