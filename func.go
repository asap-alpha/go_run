package main

import (
	"fmt"
	"time"
)

const hours = 24

func main() {
	fmt.Println("we are here")

	t := time.Now()
	switch t.Hour() {
	case hours:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
