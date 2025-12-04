package main

import "os"

func main5() {
	files, err := OSReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		println(file)
	}
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
