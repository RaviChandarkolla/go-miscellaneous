package main

import "fmt"

func main() {

	// Declare a variable of an anonymous type set to its zero value
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	fmt.Printf("%+v\n", e1)

	// Declare  a variable of an anonymous type and init using a struct literal
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Printf("%+v\n", e2)
}
