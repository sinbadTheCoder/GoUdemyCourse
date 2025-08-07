package main

import (
	"fmt"

	"example.com/array/arrays"
	"example.com/array/maps"
)

func main() {
	//
	// Arrays
	//
	arrays.Main()

	//
	// Maps
	//
	maps.Main()

	//
	// Make
	//
	userNames := make([]string, 2, 5)

	userNames[0] = "Moses"
	userNames[1] = "Joshua"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Mark")

	fmt.Println(userNames)
}
