package main

import "fmt"

func main1() {
	coursesInPg := []string{
		"Docker & Kubernetes",
		"Docker networking",
		"Kubernetes Deep Dive"}

	coursesCompleted := []string{
		"Docker Deep Dive",
		"Kubernetes Deep Dive"}

	for i, data := range coursesInPg {
		fmt.Println(i)
		fmt.Println(data)

		for _, j := range coursesCompleted {
			if data == j {
				fmt.Println("\n Hey we found a clash.", data , "is in both lists.")
			}
		}
	}
}
