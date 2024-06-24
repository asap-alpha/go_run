package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	const serviceData = "eyJTZXJ2aWNlUmVxdWVzdElkIjoiYTE1MjJiYzQzOCIsIkl0ZW1JZCI6IjNkMTNkMTlhYmQzNjM4NCIsIkNoYW5uZWwiOiJ2b2RhZm9uZSIsIk9yZGVySWQiOiJmNjMzYmE4ZTMxZCJ9"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Printf("%r", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		_ = fmt.Sprintf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("let's start here")
	for i, w := 0, 0; i < len(serviceData); i += w {
		runeValue, width := utf8.DecodeRuneInString(serviceData[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRuneValue(runeValue)
	}

}
func examineRuneValue(r rune) {
	if r == 'r' {
		fmt.Println("found tree")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
