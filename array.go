package main

import "fmt"

// Array in go has immutable size (cannot be expanded)
// len(array) to get size of array
func mainArray() {
	indexing()
	declareInitialize()
	declareUnknownSize()
}

func indexing() {
	var names [3]string
	names[0] = "dioxide"
	names[1] = "monoxide"
	names[2] = "peroxide"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
}

func declareInitialize() {
	var values = [3]string{
		"cosmicdioxide",
		"cosmicmonoxide",
		"cosmicperoxide",
	}

	fmt.Println(values)
}

func declareUnknownSize() {
	// after declaration the size will stay that way
	var names = [...]string{
		"dioxide",
		"peroxide",
		"stronk",
	}

	fmt.Println(names)
}
