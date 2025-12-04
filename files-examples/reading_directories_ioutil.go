package main

import (
	"io/ioutil"
	"log"
)

func main4() {
	files, err := IOReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		println(file)
	}
}

func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
