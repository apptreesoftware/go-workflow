package step

import (
	"fmt"
	"os"
)

func Run() {
	stepEnv := GetEnvironment()
	stepId := StepNameAndVersion(stepEnv)

	s := GetStep(stepId)
	if s != nil {
		s.Execute()
	} else {
		fmt.Printf("Step not found in package: %s", stepId)
		os.Exit(1)
	}
}
