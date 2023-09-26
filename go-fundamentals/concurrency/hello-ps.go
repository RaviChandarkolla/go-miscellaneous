package main

import (
	"fmt"
	"sync"
	"time"
)

// func main here is also a go routine, it's the main one. As soon as it exits the whole program does and the main
// program exists before all other go routines are executed.
// using waitGrp, go-routines have a way to tell main when they are finished
func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("hello")
	}()

	// sticking an empty set of operands make it self-executing
	go func() {
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()
	waitGrp.Wait()
}
