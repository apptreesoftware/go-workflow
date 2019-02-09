package main

import (
	. "github.com/apptreesoftware/go-workflow"
	"gopkg.in/yaml.v2"
)

func main() {
	input := map[string]interface{}{}

	BindInputs(&input)
	b, err := yaml.Marshal(input)
	if err != nil {
		ReportError(err)
	}
	println(string(b))
	SetOutput(nil)
}
