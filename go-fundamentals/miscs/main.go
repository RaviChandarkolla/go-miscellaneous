package main

import "fmt"

type Sample struct {
	value1 string
	value2 string
}

func main() {
	val1 := &Sample{value1: "hello", value2: "world"}
	val2 := Sample{value1: "bza", value2: "hyd"}

	passByReference(val1)
	passByValue(val2)

	fmt.Printf("val1 = %s", *val1)
	println()
	fmt.Printf("val2 = %s", val2)
}

func passByReference(s *Sample) {
	s.value1 = "Modified"
}

func passByValue(s Sample) {
	s.value1 = "Modified"
}
