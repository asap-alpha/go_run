package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	assignValueToRect := rect{height: 20, width: 10}

	fmt.Println("area =>", assignValueToRect.area())
	fmt.Println("perimeter =>", assignValueToRect.perim())

	newValue := &assignValueToRect

	newValue.height = 10
	newValue.width = 10
	fmt.Println("new area value")
	fmt.Println("area =>", newValue.area())
	fmt.Println("perimeter =>", newValue.perim())

}
