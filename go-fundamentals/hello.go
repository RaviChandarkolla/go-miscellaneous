package main

import "fmt"

func main() {
	fmt.Println("Hello main")
	var tmp *string

	print(tmp)
	if tmp == nil {
		print("yes")
	}
}
