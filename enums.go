package main

func main() {

	// enum 
	type enumType int

	const (
		area = iota
		perim
	)

	var geometryType = map[enumType]string{
		area:  "area",
		perim: "perim",
	}

	func (s enumType ) String() string{
		
		
	}
}
