package main

import "fmt"

// Will still run despite error if run by file name
func mainDataType() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Satu koma satu = ", 1.1)
	fmt.Println("Bool true = ", true)
	fmt.Println("Bool false = ", false)
	printString()
}

func printString() {
	fmt.Println("Cosmicdioxide")
	fmt.Println("Cosmicmonoxide")
	fmt.Println("String lenght: ", len("cosmicdioxide"))
	fmt.Println("First char: ", "cosmic"[0])
}
