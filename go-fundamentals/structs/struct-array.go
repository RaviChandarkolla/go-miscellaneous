package main

type Rectangle1 struct {
	Top    int
	Left   int
	Width  int
	Height int
	UpDown int
}

func main() {
	var arr []Rectangle1
	println(len(arr))
	if arr == nil {
		println("array is empty")
	}
}
