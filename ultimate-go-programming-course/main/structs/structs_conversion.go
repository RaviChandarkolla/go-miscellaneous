package main

import "fmt"

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	var a alice
	var b bill

	// this is implicit conversion if types are identical. this may work in some other languages not in go
	//b = a
	fmt.Println(a)
	fmt.Println(b)

	// this is explicit conversion
	b = bill(a)
	fmt.Println(a)
	fmt.Println(b)

	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// implicit conversion can happen b/w name type and literal type. This type of assignment with literal values you
	// will see it in functions. A function in go is a literal value, it is not a named type
	b = e3
	fmt.Println(a)
	fmt.Println(b)

}
