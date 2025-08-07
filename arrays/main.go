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

	makeSlice()
	makeMap()
}

func makeSlice() {
	//create array
	namesArray := [4]string{}

	namesArray[0] = "Moses"
	namesArray[1] = "Joshua"
	namesArray[2] = "John"
	namesArray[3] = "Mark"

	fmt.Println("Usernames array:", namesArray)

	//create slice
	namesSlice := []string{"Moses", "Joshua"}

	namesSlice = append(namesSlice, "John")
	namesSlice = append(namesSlice, "Mark")

	fmt.Println("Usernames array:", namesSlice)

	//create a slice through make
	namesSlice = make([]string, 2, 5)

	namesSlice[0] = "Moses"
	namesSlice[1] = "Joshua"
	namesSlice = append(namesSlice, "John")
	namesSlice = append(namesSlice, "Mark")

	fmt.Println("Usernames array through Make slice", namesSlice)
}

func makeMap() {
	languagePopularityMap := map[string]float64{}
	languagePopularityMap["php"] = 4.1
	languagePopularityMap["go"] = 4.3
	languagePopularityMap["java"] = 4.5
	languagePopularityMap["c"] = 3.2
	languagePopularityMap["rust"] = 3.9

	fmt.Println("Language Popularities:", languagePopularityMap)

	languagePopularityMap = make(map[string]float64, 5)
	languagePopularityMap["php"] = 4.1
	languagePopularityMap["go"] = 4.3
	languagePopularityMap["java"] = 4.5
	languagePopularityMap["c"] = 3.2
	languagePopularityMap["rust"] = 3.9

	fmt.Println("Language Popularities (from make):", languagePopularityMap)
}
