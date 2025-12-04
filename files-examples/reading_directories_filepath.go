package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main8() {
	// using filepath.Walk
	files, err := FilePathWalkDir("./")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
