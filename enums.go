package main

import "fmt"

// enum
type enumType int

const (
	area enumType = iota
	perim
)

//var geometryType = map[enumType]string{
//	area:  "area",
//	perim: "perim",
//}

func String(s enumType) string {
	switch s {
	case area:
		return "Area"
	case perim:
		return "Perimeter"
	default:
		panic("no value assigned, we break out.")

	}

}

func main() {

	fmt.Println(String(area))

}
