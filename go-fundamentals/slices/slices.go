package main

import "fmt"

func main1() {
	// type, length and capacity
	courses := make([]string, 5, 10)
	courses[0] = "Docker & Kubernetes"
	courses[1] = "Getting started with Docker"
	courses[2] = "Getting started with Kubernetes"

	// this below code fails
	//courses[11] = "Testing"
	//fmt.Println(courses[11])

	fmt.Printf("Length of slice is %d and capacity if %d\n\n", len(courses), cap(courses))

	fmt.Println(courses)

	for _, i := range courses {
		fmt.Println(i)
	}
}
