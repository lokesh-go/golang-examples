package main

import (
	"log"

	pkg "github.com/lokesh-go/golang-examples/databinding/pkg"
	printPkg "github.com/lokesh-go/golang-examples/print/pkg"
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
	bytes, err := pkg.JSONMarshal(testMap)
	if err != nil {
		log.Fatal(err)
	}

	// Now unmarshal json encoded data into struct
	tstruct := testStruct{}
	err = pkg.JSONUnmarshal(bytes, &tstruct)
	if err != nil {
		log.Fatal(err)
	}

	// Prints
	// Basically we had used json marshal & unmarshal method to bind the data
	printPkg.PrintWithNewLine(tstruct)

	// Now data binding using mapstructure
	// Basically it is a third party data binding provider
	// which is binding data one to another
	output := testStruct{}
	err = pkg.MapStructureDecode(testMap, &output)
	if err != nil {
		log.Fatal(err)
	}

	// Prints
	printPkg.Print(output)
}
