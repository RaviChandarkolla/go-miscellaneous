package main

import (
	"encoding/json" //JSON package from the standard library
	"fmt"
	"os"
)

// Rectangle represents a rectangle shape
type Rectangle struct {
	Top    int
	Left   int
	Width  int
	Height int
	UpDown int
}

type RectangleCustomized struct {
	Top    int `json:"top,omitempty"`
	Left   int `json:"left,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"-"` //will never be encoded
	UpDown int `json:"up-down,omitempty"`
}

func main3() {
	//create an instance of a Rectangle
	r := &Rectangle{10, 20, 30, 40, 50}

	//create a new JSON encoder over stdout
	//and encode the struct into JSON
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(r); err != nil {
		fmt.Printf("error encoding struct into JSON: %v\n", err)
	}

	//create an instance of a Rectangle
	rc := RectangleCustomized{10, 20, 30, 40, 50}

	//marshal into JSON buffer
	buffer, err := json.Marshal(rc)
	if err != nil {
		fmt.Printf("error marshaling JSON: %v\n", err)
	}

	fmt.Println(string(buffer))

	//create an empty instance of a Rectangle
	//to receive the decoded JSON
	//note that r is a pointer to a Rectangle
	dr := &Rectangle{}

	//create a new decoder over stdin
	//and decode the JSON into the struct
	// {"Top":50,"Left":60,"Width":70,"Height":80} - give this as input for stdin

	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(dr); err != nil {
		fmt.Printf("error decoding JSON: %v\n", err)
	}

	fmt.Println(dr)
}
