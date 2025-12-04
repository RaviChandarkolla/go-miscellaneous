package main

import (
	"encoding/json"
	"fmt"
)

type eg struct {
	val1 interface{}
	val3 *string
}

type eg1 struct {
	val1 interface{}
	val3 *string
}

func main() {
	tmp1 := eg{
		val1: 10,
	}
	var tmp2 eg
	egContent, err := json.Marshal(tmp1)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(egContent, tmp2)
	if err != nil {
		panic(err)
	}

	fmt.Sprintf("%v", tmp2)
}
