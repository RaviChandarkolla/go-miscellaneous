package main

func main() {
	count := 10

	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// we will make a copy of the data and is sent in request parameters
	increment(count)
	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")
}

func increment(inc int) {
	inc++
	println("count:\tValue of[", inc, "]\tAddr of[", &inc, "]")
}
