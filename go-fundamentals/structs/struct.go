package main

import "fmt"

func main2() {
	// this is just defining a type
	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	//var gettingStarted courseMeta

	//// using the new keyword yields a pointer
	//gettingStarted1 := new(courseMeta)

	gettingStartedWithK8s := courseMeta{
		author: "Ravi",
		level:  "Basic",
		rating: 5,
	}
	fmt.Println(gettingStartedWithK8s)

	fmt.Println("Author of course is", gettingStartedWithK8s.author)
	gettingStartedWithK8s.rating = 1.2
	fmt.Println("New rating of course is", gettingStartedWithK8s.rating)
}
