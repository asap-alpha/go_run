package main

import "fmt"

func main() {
	fmt.Println("we are here to learn Pointers.")
	var value int = 1

	noPointer(value)
	usingPointer(&value)

}

func noPointer(noPoint int) int {

	result := 2 * noPoint
	fmt.Println(result)
	return result
}

func usingPointer(usePoint *int) int {

	result := *usePoint * 5
	fmt.Println(result)
	return result
}
