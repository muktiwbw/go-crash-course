package main

import "fmt"

// Slice is like array, but the size is dynamic/mutable
// len(slice) to get length of slice
// cap(slice) to get slice's capacity
// append(slice, newData) to add data to slice
// make([]type, length, capacity) create a slice
// copy(destination, source) copy a slice
func mainSlice() {
	sliceFromArray()
}

func sliceFromArray() {
	var values = [...]int8{
		10,
		11,
		12,
		13,
		14,
		15,
		16,
	}

	// from index 2 to 3 (exclusive end)
	firstSlice := values[2:4]
	fmt.Println(firstSlice)

	var secondSlice []int8 = values[3:]
	var thirdSlice []int8 = values[:5]
	var fourthSlice = values[:]

	fmt.Println(secondSlice)
	fmt.Println(thirdSlice)
	fmt.Println(fourthSlice)

}
