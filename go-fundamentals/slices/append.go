package main

import "fmt"

func main3() {
	mySlice := make([]int, 1, 4)

	fmt.Printf("Length starts out as %d with a capacity of %d \n", len(mySlice), cap(mySlice))

	for i := 1; i <= 10; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Slice length is %d but capacity is %d \n", len(mySlice), cap(mySlice))
	}

	mySlice1 := make([]int, 1, 4)
	fmt.Printf("Length starts out as %d with a capacity of %d \n", len(mySlice1), cap(mySlice1))

	newSlice := []int{10, 20, 30, 40}
	mySlice1 = append(mySlice1, newSlice...)
	fmt.Printf("Length starts out as %d with a capacity of %d \n", len(mySlice1), cap(mySlice1))
}
