package main

import (
	"fmt"
	maps2 "maps"
	"slices"
)

func main() {

	fmt.Println("we are leveling up")
	//slice()
	//sliceLength()
	//maps()
	rangee()
}

func arrayy() {

	var myArray [4]int

	myArray[0] = 1
	myArray[1] = 11
	myArray[2] = 111
	myArray[3] = 1111

	for i := range myArray {
		fmt.Printf("index %d the values %d \n", i, myArray[i])
	}

}

func arrayTwo() {
	fmt.Println(">>>>>>>>>>>>more level up")

	newArray := [5]int{10, 12, 13, 14, 15}

	newArray[4] = 100

	for v := range newArray {
		fmt.Printf("index %d values %d \n", v, newArray[v])
	}

}

func SpreadArray() {

	fmt.Println(">>>>>>>>>>>using spread>>>>>")

	spreadArray := [...]int{1, 2, 3, 10: 10}
	spreadArray[9] = 9
	spreadArray[4] = 8
	for i := 0; i < len(spreadArray); i += 1 {

		if spreadArray[i]%2 != 0 {
			fmt.Println("Divisible by 2", spreadArray[i])
			continue
		}

		fmt.Printf("index %d, spread values %d \n", i, spreadArray[i])
	}
}

func twoDArray() {
	var twoD [3][2]int

	for a := 0; a < 3; a += 1 {
		for b := 0; b < 2; b += 1 {
			fmt.Printf("b %d \n", b)
			twoD[a][b] = a + b

		}
		fmt.Printf("a %d \n", a)
	}
	fmt.Println(twoD)
}

func slice() {
	s := make([]string, 10)
	s[0] = "Ghana"
	s[1] = "Togo"
	s[2] = "Cote-D'voire"
	s[3] = "Burkina-Faso"
	s[4] = "Niger"
	s[5] = "Nigeria"

	s = append(s, "Morocco")

	for a := range s {
		fmt.Println(s[a])
	}

	l := s[10:]
	fmt.Println("value of slices", l)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)
}

func sliceLength() {
	length1 := []string{"Ama", "Kofi", "yaw"}

	length2 := []string{"Ama", "Kofi", "yaw"}

	if slices.Equal(length2, length1) {
		fmt.Printf("true")
	}

	twoSlice := make([][]int, 3)

	for a := 0; a < 3; a += 1 {
		length := a + 1
		twoSlice[a] = make([]int, length)
		fmt.Printf("values of a %d \n", twoSlice[a])

		for b := 0; b < length; b++ {
			twoSlice[a][b] = a + b
			fmt.Printf("values of b %d \n", twoSlice[b])
		}
		fmt.Printf("values of a second side: %d \n", twoSlice[a])

	}

	fmt.Printf("two Slice: %d \n", twoSlice)
}

func maps() {
	m := make(map[string]int)

	m["k1"] = 20
	m["k2"] = 21

	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println(len(m))

	delete(m, "k1")
	fmt.Println(len(m))
	clear(m)
	fmt.Println(len(m))

	a := map[string]int{"k3": 2, "K4": 23, "k5": 90}

	maps2.Copy(m, a)
	fmt.Println("newly assigned values: ", m)

	theSame := maps2.Equal(m, a)

	switch theSame {
	case true:
		fmt.Println("both maps are equal")
	default:
		fmt.Println("Maps inconsistency")
	}
}

func rangee() {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0

	for _, num := range nums {
		sum += num

		fmt.Println("sum", sum)
	}

	kvs := map[string]int{"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5, "k6": 6}
	sum2 := 0

	for k, v := range kvs {

		fmt.Println("key: value:", k, v)

		sum2 += v

		fmt.Println("summation: ", sum2)
	}

}
