package main

import (
	"log"

	dataBindingPkg "golang-examples/databinding/pkg"
)

type testStruct struct {
	Language string  `json:"language"`
	Version  float64 `json:"version"`
}

// Package testing file
func main() {
	// Defines map & assign value
	testMap := map[string]interface{}{
		"language": "golang",
		"version":  1.16,
	}

	// Lets bind testMap data into struct
	// Firstly we have to do json encode
	bytes, err := dataBindingPkg.JSONMarshal(testMap)
	if err != nil {
		log.Fatal(err)
	}

	// Now unmarshal json encoded data into struct
	tstruct := testStruct{}
	err = dataBindingPkg.JSONUnmarshal(bytes, &tstruct)
	if err != nil {
		log.Fatal(err)
	}

}
