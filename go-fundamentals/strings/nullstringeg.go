package main

import (
	"gopkg.in/guregu/null.v4"
)

func main1() {
	//var str null.String
	//
	//var str1 string

	//println(str.Valid)
	//println(str.IsZero())
	//println(str.String)
	//println(len(str.String))
	//output := null.StringFrom(str1)
	//
	//fmt.Println(output)
	//fmt.Println(output.Valid)
	//fmt.Println(output.IsZero())
	//if output.String == "" {
	//	println("not equal")
	//}
	//
	//if str.String == "" {
	//	println("not equal 2")
	//}
	//
	//str3 := null.StringFrom("")
	//if str3.String == "" {
	//	println("not equal 3")
	//}

	var str4 *string
	output := null.StringFrom(*str4)
	println(output.String)
}
