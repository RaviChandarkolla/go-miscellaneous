package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Length starts out as %d with a capacity of %d \n", len(mySlice), cap(mySlice))

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Printf("Length starts out as %d with a capacity of %d \n", len(mySlice), cap(mySlice))

}
