package main

import "fmt"

func main() {
	argAsPointer()
}

type Address struct {
	City, Province, Country string
}

func learnPointer() {
	add1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	add2 := add1
	add2.City = "Malang"

	fmt.Println(add1)
	fmt.Println(add2)

	// Value from add1 didn't change because add1 data was passed by value to add2
}

func passByReference() {
	add1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	add2 := &add1
	add2.City = "Malang"

	fmt.Println(add1)
	fmt.Println(add2)

	// Value of add1 was changed because add2 was changed
	// add2 was a pointer to add1
}

func asteriskOperator() {
	add1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	add2 := &add1
	add2.City = "Malang"

	*add2 = Address{"Solo", "Jawa Tengah", "Indonesia"}

	fmt.Println(add1)
	fmt.Println(add2)
}

func argAsPointer() {
	add1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println("Before change:", add1)
	changeCityToJember(&add1)
	fmt.Println("After change:", add1)
}

func changeCityToJember(address *Address) {
	address.City = "Jember"
}
