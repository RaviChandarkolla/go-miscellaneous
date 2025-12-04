package main

import "fmt"

type CustomPrintFn interface {
	printFn() string
}

type Book struct {
	Title  string
	Author string
}

func (b Book) printFn() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.Title, b.Author)
}

type Count int

func (count Count) printFn() string {
	return fmt.Sprintf("Count: %d", count)
}

func WriteLog(cpf CustomPrintFn) {
	fmt.Println(cpf.printFn())
}

func main() {
	// Initialize a Count object and pass it to WriteLog().
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	book.printFn()
	//WriteLog(book)
	//
	//// Initialize a Count object and pass it to WriteLog().
	//count := Count(3)
	//WriteLog(count)
}
