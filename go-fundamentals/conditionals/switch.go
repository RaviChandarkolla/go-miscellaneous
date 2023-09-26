package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {
	switch "kubernetes" {
	case "kubernetes":
		fmt.Println("Case 1. kubernetes is hit")
		fallthrough
	case "Kubernetes":
		fmt.Println("Case 2. Kubernetes with capital K is hit")
		fallthrough
	case "k8s":
		fmt.Println("Case 3. k8s is hit")
		fallthrough
	default:
		fmt.Println("Case 4. Hit the default")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got an even number - ", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got odd number - ", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)

}
