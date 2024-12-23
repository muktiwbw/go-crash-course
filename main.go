package main

import "fmt"

func main() {
	tryDefer()
	tryPanic(true)
}

func logging() {
	fmt.Println("Function has been executed!")
}

// Any function called with defer keyword would be executed after everything within the same scope
// is done executed.
func tryDefer() {
	defer logging()
	fmt.Println("Doing something...")
}

// Use recover() inside of a function that will be called using defer
func recoverPanic() {
	message := recover()
	fmt.Println("Some error occured with message: ", message)
}

/*
	Test
*/
// Error that stops the program entirely
func tryPanic(isError bool) {
	defer recoverPanic()
	if isError {
		panic("Error ocurred!")
	}
}
