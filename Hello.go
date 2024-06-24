package main

import (
	"fmt"
)

func main() {
	//const pi float64 = 31 / 2
	//name := "Abdallah"
	//var homeAddress string
	//fmt.Println("hello GO!", name)
	//homeAddress = "Afere"
	//fmt.Println("home Address", homeAddress)
	//fmt.Println("pi value", pi)

	//i := 0
	//for i < 3 {
	//	i += 1
	//	fmt.Println("value", i)
	//}
	//
	//for j := 0; j < 10; j++ {
	//	fmt.Println("value for j:", j)
	//	if j == 5 {
	//		break
	//	}
	//}
	//
	//for v := range 5 {
	//	fmt.Println("value for range", v)
	//}
	//
	//var value int
	//fmt.Printf("Input value %v \n", value)
	//fmt.Scan(&value)
	//
	//if value == 5 {
	//	fmt.Println("value is ", value)
	//	return
	//}
	//{
	//	fmt.Printf("value entered was %v \n", value)
	//}

	//hours

	//tim := time.Now().Truncate(time.Minute)
	//
	//fmt.Println(tim)
	//
	//hour := time.Now()
	//
	//switch {
	//case hour.Hour() < 12:
	//	fmt.Printf("the time is less %d:00Hr \n", hour.Hour())
	//	break
	//case hour.Hour() > 12:
	//	fmt.Printf("the time is more %d:00Hr \n", hour.Hour())
	//	break
	//}
	////weekend
	//
	//switch time.Now().Weekday() {
	//case time.Saturday, time.Sunday:
	//	fmt.Printf("we are weekend ", time.Now().Weekday().String())
	//	break
	//default:
	//	fmt.Printf("we are in weekdays ", time.Now().Weekday().String())
	//}

	//take input and check switch
	//var input int
	//fmt.Printf("enter a number")
	//fmt.Scan(&input)
	//
	//switch input % 2 {
	//case 0:
	//	fmt.Printf("%d is even number:", input)
	//	break
	//default:
	//	fmt.Printf("%d is old number:", input)
	//}

	//whatIsIt := func(i interface{}) {
	//
	//	switch t := i.(type) {
	//	case int:
	//		fmt.Printf("This is integer %d \n", t)
	//	case bool:
	//		fmt.Printf("This is boolean %d \n", t)
	//	case string:
	//		fmt.Printf("This is string %d \n", t)
	//	default:
	//
	//		fmt.Printf("this is none", t)
	//	}
	//}

	//whatIsIt(12.33)

	//array

	var arr [5]int

	arr[3] = 10
	fmt.Println("this is the array", arr)
	fmt.Println("this prints the position", arr[3])

	ars := [5]int{11, 12, 13, 14, 15}

	fmt.Printf("contains", ars)
	fmt.Println("this contains the length of the array", len(ars))

	arDot := [...]int{20, 22, 4: 23, 24}
	fmt.Println(arDot)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
			fmt.Println(twoD)
		}
	}

	fmt.Println("second", twoD)
}
