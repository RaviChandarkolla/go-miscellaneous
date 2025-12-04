package main

import (
	"fmt"
	"log"
	"os"
)

func main2() {
	// normal listing of files present in directory
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}

	// using filepath.Walk
	files, err := FilePathWalkDir("./")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
