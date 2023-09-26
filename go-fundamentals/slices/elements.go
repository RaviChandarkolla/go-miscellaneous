package main

import "fmt"

func main2() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice)
	fmt.Println(mySlice[4])

	mySlice[1] = 0
	fmt.Println(mySlice)

	// index 2 is included but 5 is not
	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)

	sliceFromIndex0To5 := mySlice[:5]
	fmt.Println(sliceFromIndex0To5)

	sliceFromIndex5ToEnd:= mySlice[5:]
	fmt.Println(sliceFromIndex5ToEnd)
}
