package main

import (
	. "github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	input := Input{}
	//Use the light weight sdk to bind the inputs to your struct
	BindInputs(&input)

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
