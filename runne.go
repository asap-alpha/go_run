package main

import "fmt"

func main() {
	const s = "สวัสดี"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
}
