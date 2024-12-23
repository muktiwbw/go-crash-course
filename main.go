package main

import (
	"fmt"
	"strings"
)

func main() {
	simpleCondition()
	shortStatement()
	switchExpression()
	forLoop()
	shorterForLoop()
	forRange()
	fmt.Println(addition(10, 12))
	fmt.Println(multipleUppercase("cosmic", "peroxide"))
	fmt.Println(namedReturnValues())
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	numbers := []int{1, 2, 3}
	fmt.Println(sumAll(numbers...))

	variadicSumAll := sumAll
	fmt.Println("Called from variadicSumAll: ", variadicSumAll(numbers...))
	fmt.Println("4! =", factorial(4))
	fmt.Println("5! =", factorial(5))

}

func simpleCondition() {
	name := "cosmic"

	if name == "cosmic" {
		fmt.Println("Hello cosmic")
	} else {
		fmt.Println("Who?")
	}
}

func shortStatement() {
	name := "cosmicdioxide"

	if length := len(name); length > 5 {
		fmt.Println("Too long")
	} else {
		fmt.Println("Okay!")
	}
}

func switchExpression() {
	name := "cosmicdioxide"

	switch name {
	case "cosmicdioxide":
		fmt.Println("Hello cosmic")
	case "cosmicperoxide":
		fmt.Println("Huh?")
	default:
		fmt.Println("Bruh")
	}

	// dumb code
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Yo!")
	default:
		fmt.Println("Good!")
	}
}

func forLoop() {
	counter := 0

	for counter < 10 {
		fmt.Println("Counter:", counter+1)
		counter++
	}
}

func shorterForLoop() {
	for counter := 0; counter < 10; counter++ {
		fmt.Println("Shorter counter: ", counter+1)
	}
}

func forRange() {
	listOfEntry := map[string]string{
		"firstName": "cosmic",
		"lastName":  "dioxide",
		"age":       "99",
	}

	for key, val := range listOfEntry {
		fmt.Println("Key:", key, "Val:", val)
	}

	listOfDays := []string{
		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
	}

	for i, day := range listOfDays {
		fmt.Println("Index: ", i, "Day: ", day)
	}

	// Without index
	for _, day := range listOfDays {
		fmt.Println("Day: ", day)
	}
}

func addition(a int, b int) int {
	return a + b
}

func multipleUppercase(name1 string, name2 string) (string, string) {
	return strings.ToUpper(name1), strings.ToUpper(name2)
}

func namedReturnValues() (firstName, lastName string) {
	firstName = "cosmic"
	lastName = "monoxide"

	return firstName, lastName
}

// Variadic functions
// variadic args should be at the very end
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func factorial(number int) int {
	if number == 1 {
		return number
	}

	return number * factorial(number-1)
}
