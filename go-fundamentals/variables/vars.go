package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name, course = "Ravi", "Go Fundamentals"
	module, clip = "4", 2
	name1, course1 string
	module1, clip1 int
)

func main1() {
	fmt.Println("Name and course are set to ", name, " and ", course, ".")
	fmt.Println("Module and clip are set to ", module, " and ", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}

	fmt.Println("Memory address of *course* variable is", &course)

	var ptr *string = &course
	fmt.Println("pointing course variable at address,", ptr," which holds this value", *ptr)
}
