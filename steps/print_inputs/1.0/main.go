package main

import (
	. "github.com/apptreesoftware/go-workflow/pkg/step"
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
