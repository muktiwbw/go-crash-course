package main

import "fmt"

// All veriables should be used or the compiler will fail the during build
// Unlike variables, constants can left unused
func main() {
	var firstName string
	var lastName string = "dioxide"

	firstName = "cosmic"
	fmt.Println("name: ", firstName, lastName)

	initializeVar()
	multipleInitializeVar()
	multipleConstants()
}

func initializeVar() {
	firstName := "cosmic"
	lastName := "monoxide"

	fmt.Println("Second name:", firstName, lastName)
}

func multipleInitializeVar() {
	var (
		firstName = "cosmic"
		lastName  = "peroxide"
	)

	fmt.Println("Third name:", firstName, lastName)
}

func multipleConstants() {
	const (
		firstName = "cosmic"
		lastName  = "desert"
	)

	fmt.Println("Fourth name:", firstName, lastName)
}
