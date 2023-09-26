package main

import (
	"fmt"
)

func main1() {
	dddLengthMins := 275
	cawLengthMins := 30

	if dddLengthMins > cawLengthMins {
		fmt.Println("Docker deep dive is longer than containers on AWS wavelength")
	} else if dddLengthMins > cawLengthMins {
		fmt.Println("Docker deep dive is same length as AWS wavelength")
	} else {
		fmt.Println("containers on AWS wavelength is longer than Docker deep dive")
	}

	if tmp1, tmp2 := 275, 30; tmp1 > tmp2 {
		fmt.Println("tmp1 is greater than tmp2")
	}
}
