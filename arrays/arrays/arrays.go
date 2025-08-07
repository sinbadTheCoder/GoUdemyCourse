package arrays

import (
	"fmt"
	"time"
)

type Course struct {
	CourseName     string
	InstructorName string
	Curriculum     string
	Students       []string
	LengthMinutes  int
	StartDate      time.Time
}

func listCourses() {
	courses := []Course{}
	firstCourse := Course{
		CourseName:     "Defense Against the Dark Arts",
		InstructorName: "Professor Lupin",
		Curriculum:     "Animangus and Warewolves",
		Students:       []string{"Harry Potter", "Ron Weasly", "Hermione Granger"},
		LengthMinutes:  90,
		StartDate:      time.Now(),
	}

	courses = append(courses, firstCourse)

	secondCourse := Course{
		CourseName:     "Potions",
		InstructorName: "Professor Snape",
		Curriculum:     "Tom Riddle's Diary",
		Students:       []string{"Harry Potter", "Ron Weasly", "Hermione Granger"},
		LengthMinutes:  60,
		StartDate:      time.Now(),
	}

	courses = append(courses, secondCourse)

	thirdCourse := Course{
		CourseName:     "Herbology",
		InstructorName: "Professor McGonigal",
		Curriculum:     "Mandrakes and other screaming plants",
		Students:       []string{"Harry Potter", "Ron Weasly", "Hermione Granger"},
		LengthMinutes:  60,
		StartDate:      time.Now(),
	}

	courses = append(courses, thirdCourse)

	fmt.Println("~~~~~~~~~~~~~~~ Courses ~~~~~~~~~~~~~~~")
	for _, course := range courses {
		fmt.Printf("Course Name: %s\n", course.CourseName)
		fmt.Printf("\tInstructor Name: %s\n", course.InstructorName)
		fmt.Printf("\tCurriculum: %s\n", course.Curriculum)
		fmt.Println("\tStudents: ", course.Students)
		fmt.Printf("\tCourse Length (Minutes): %d\n", course.LengthMinutes)
		fmt.Printf("\tStart Date: %s\n", course.StartDate)
		fmt.Print("\n")
	}

}

func Main() {
	listCourses()
	slicesBasics()
}

func slicesBasics() {
	// setup pricess for products
	prices := [4]float64{14.99, 198.99, 54.99, 20}

	// now provide product names
	var productNames [4]string = [4]string{"A Book"}

	// now assign name to third element of array
	productNames[2] = "A Carpet"

	// printing prices array works in Printf
	fmt.Printf("%.1f\n", prices)
	// print product names
	fmt.Printf("%s\n", productNames)

	fmt.Println("Full price list:", prices)
	fmt.Println("2nd element: ", prices[1])
	fmt.Println("2nd through 3rd elements:", prices[1:3])
	fmt.Println("1st through 3rd elements:", prices[:3])
	fmt.Println("1st through 3rd elements:", prices[0:3])
	fmt.Println("2nd element through end:", prices[1:])

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Println("Highlighted Prices:", highlightedPrices)

	//replace element of orig array by modifying slice (increasing by 1 dollar)
	highlightedPrices[0] += 1
	fmt.Println("Modified price:", prices[1])

	// Merge two slices together
	fmt.Println(prices)
	pricesSlice := prices[:]
	newlyDefinedPrices := []float64{9.99, 10.99, 11.99}
	pricesSlice = append(pricesSlice, newlyDefinedPrices...)
	fmt.Println("Merged list of prices:", pricesSlice)
}
