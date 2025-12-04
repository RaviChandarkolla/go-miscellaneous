package main

import "strings"

func main() {
	str := "Welcome:to:GeeksforGeeks"
	listStr := str[:strings.LastIndex(str, ":")]
	println(listStr)
}
