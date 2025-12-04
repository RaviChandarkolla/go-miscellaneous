package main

import (
	"fmt"
	"unsafe"
)

type example struct {
	flag    bool
	counter int16
	pi      float32
	flag1   bool
}

type optimised struct {
	flag  bool
	pi    float64
	flag1 bool
}

func main() {
	// declare a variable of type example and set to its zero value
	var e1 example
	fmt.Printf("%+v\n", e1)
	fmt.Sprintf("declared example type size = %d", unsafe.Sizeof(e1))

	op := optimised{}
	fmt.Println(unsafe.Sizeof(op))
	fmt.Sprintf("initialised  optimised type size = %d", unsafe.Sizeof(op))

	// Declare a variable of type example and init using a struct literal - also called as literal construction
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14,
	}
	fmt.Sprintf("Declared and initialised example type size = %d", unsafe.Sizeof(e2))

	fmt.Println("flag = ", e2.flag)
	fmt.Println("counter = ", e2.counter)
	fmt.Println("pi = ", e2.pi)
}
