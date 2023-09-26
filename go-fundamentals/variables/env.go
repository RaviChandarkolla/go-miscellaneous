package main

import (
	"fmt"
	"os"
)

func main() {
	// prints all environment variables on our system
	fmt.Println(os.Environ())

	for  _, env := range os.Environ() {
		fmt.Println(env)
	}

	name3 := os.Getenv("USER")
	println(name3)
}