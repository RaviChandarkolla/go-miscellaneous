package main

import "fmt"

func main() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
	}

	for mapKey, mapVal := range testMap {
		fmt.Printf("Key is :%v Value is: %v\n", mapKey, mapVal)
	}

	fmt.Println(testMap)
	fmt.Println(testMap["C"])

	testMap["A"] = 100
	fmt.Println(testMap)

	testMap["J"] = 200
	fmt.Println(testMap)

	delete(testMap, "J")
	fmt.Println(testMap)
}
