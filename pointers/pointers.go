package main

import "fmt"

// Demonstrates how to avoid copying values of variables through
// use of pointers
func main() {
	age := 32
	var agePtr *int
	fmt.Printf("Pointer address [%d]\n", agePtr)

	//this causes a segfault because value is not yet set. is nil
	/*
		_, err := fmt.Printf("Value of pointer [%d]\n", *agePtr)
		if err != nil {
			fmt.Println(err)
		}
	*/

	agePtr = &age
	fmt.Printf("Value of pointer to int variable is [%d] at address [%d]\n", *agePtr, agePtr)

	fmt.Println("Age: ", *agePtr)
	adultYears := getAdultYears(agePtr)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	fmt.Printf("Value of pointer to int variable is [%d] at address [%d]\n", *age, age)
	return *age - 18
}
