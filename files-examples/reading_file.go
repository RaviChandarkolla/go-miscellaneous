package main

import (
	"fmt"
	"os"
)

func main7() {
	contents, err := os.ReadFile("example2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}
