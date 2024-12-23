package main

import "fmt"

func main() {
	var cust Customer
	cust.Name = "cosmicdioxide"
	cust.Address = "Mars"
	cust.Age = 99

	fmt.Println(cust.Name, cust.Address, cust.Age)

	cust2 := Customer{
		Name:    "cosmicmonoxide",
		Address: "Pluto",
		Age:     192,
	}

	cust3 := Customer{"cosmicperoxide", "Neptune", 102}

	fmt.Println(cust2)
	fmt.Println(cust3)

	cust2.identify()

	typeAssertions("Hello from type assertion!")
	switchAssertions("Goodbye")
	switchAssertions(11)
	switchAssertions(nil)
}

/*
Struct name and attributes should be written as capitalized
*/
type Customer struct {
	Name, Address string
	Age           int
}

/*
Struct method
*/
func (customer Customer) identify() {
	fmt.Println("My name is", customer.Name)
}

/*
Interface
*/
type HasName interface {
	GetName() string
}

func typeAssertions(arg any) {
	fmt.Println(arg.(string))
}

func switchAssertions(arg any) {
	switch argType := arg.(type) {
	case string:
		fmt.Println("Is string", argType)
	case int:
		fmt.Println("Is int", argType)
	default:
		fmt.Println("Something else")
	}
}
