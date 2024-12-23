package main

import "fmt"

// len(map) to get lenght of map
func mainMap() {
	basicMap()
}

func basicMap() {
	person := map[string]string{
		"firstName": "cosmic",
		"lastName":  "dioxide",
	}

	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
}
