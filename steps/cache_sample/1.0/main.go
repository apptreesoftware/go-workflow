package main

import (
	. "github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	input := Input{}
	//Use the light weight sdk to bind the inputs to your struct
	BindInputs(&input)

	cache := GetCache()
	rec := MyRecord{
		Name:    "Matthew Smith",
		Address: "1234 Main st",
	}
	err := cache.PutRecord("1111", &rec, map[string]interface{}{"status": "synced"}, "default")
	ReportError(err) //If err != nil it will report it and shut down this step

	// Perform some logging to stdout.
	// This will show up in the workflow output
	println("Hello from String length counter step")

	//Perform your logic
	count := len(input.String)

	//Use the SDK to set your output easily
	SetOutput(&Output{Length: count})
}

type Input struct {
	String string
}

type Output struct {
	Length int
}

type MyRecord struct {
	Name    string
	Address string
}
