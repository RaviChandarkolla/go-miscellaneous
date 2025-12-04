package main

func main() {
	count := 10

	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// we will make a copy of the data and is sent in request parameters
	incrementPBR(&count)
	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")
}

func incrementPBR(inc *int) {
	// here inc will store the address of count value 10.
	// this is an indirect memory read and write
	*inc++
	println("count:\tValue of[", *inc, "]\tAddr of[", &inc, "]")
}

