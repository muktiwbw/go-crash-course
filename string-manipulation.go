package main

import "fmt"

func mainStringManipulation() {
	firstName := "cosmicdioxide"
	bytStrSplice := firstName[0]
	strByt := string(bytStrSplice)

	fmt.Println("first name:", firstName)
	fmt.Println("byte of first char:", bytStrSplice)
	fmt.Println("char:", strByt)
}
