package main

import "fmt"

func main() {
	var ptr *string

	fmt.Println(ptr)
	//fmt.Print(*ptr)

	var map1 map[string]interface{}
	fmt.Println(map1["check"])
}
