package main

import "fmt"

func main2() {
	name2 := "Ravi"
	course2 := "Go Fundamentals"

	fmt.Println("\n1Hi", name2, "Your current course is", course2)

	updateCourse(course2)
	fmt.Println("\n2Hi", name2, "Your current course is", course2)

	updateCourseByReference(&course2)
	fmt.Println("\n3Hi", name2, "Your current course is", course2)
}


func updateCourse(course2 string) string {
	course2 = "c++ Fundamentals"
	fmt.Println("Your current course is", course2)
	return course
}

func updateCourseByReference(course2 *string) *string {
	*course2 = "java Fundamentals"
	fmt.Println("Your current course is", *course2)
	return course2
}
