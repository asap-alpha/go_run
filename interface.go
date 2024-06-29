package main

import (
	"fmt"
	"math"
)

// interface
type Geometry interface {
	area() float64
	perimeter() float64
}

// struct
type React struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// method calculate the area and perimeter for react and struct
func (r React) area() float64 {

	fmt.Println("area of a rectangle")
	return r.width * r.height
}
func (r React) perimeter() float64 {
	result := 2*r.width + 2*r.height

	fmt.Println("perimeter of a rectangle")

	return math.Round(result)
}

func (c Circle) area() float64 {

	result := math.Pi * c.radius * c.radius
	fmt.Println("float64 value of area ==> roundedToEven", result)
	return math.RoundToEven(result)
}

func (c Circle) perimeter() float64 {
	result := 2 * (math.Pi * c.radius)
	fmt.Println("float64 value of perimeter ==> rounded", result)

	return math.Round(result)
}

// using enum to run a method
type geoEnum int

const (
	runArea geoEnum = iota
	runperimeter
)

// a function behalf as inheritor of the geometry interface
func measure(geometry Geometry, enum geoEnum) {
	fmt.Println("geometry value => ", geometry)

	switch enum {
	case runArea:
		fmt.Println(geometry.area())
		break
	case runperimeter:
		fmt.Println(geometry.perimeter())
		break
	default:
		panic("Not Found!")
	}
}

func main() {

	rectangle := React{height: 10, width: 2}
	measure(rectangle, runperimeter)

	circle := Circle{radius: 4}
	measure(circle, runperimeter)

	fmt.Println("lets run for area")
	rect := React{height: 10, width: 2}
	measure(rect, runArea)

	cir := Circle{radius: 4}
	measure(cir, runArea)
}
