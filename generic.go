package main

import "fmt"

type parameterType interface {
	int | string | float64
}

func add[T parameterType](x T, y T) T {

	return x + y
}

func main() {
	result := add(20, 3.4)

	fmt.Println(result)
}
