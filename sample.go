package main

import "fmt"

// Duplicated function in teh same package is not allowed
// func main() ---> will result in error when app is built
func helloSample() {
	fmt.Println("Hello form sample!")
}
