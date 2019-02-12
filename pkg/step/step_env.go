package step

import "os"

func GetEnvironment() *Environment {
	return &Environment{
		App:         os.Getenv("APP"),
		WorkflowId:  os.Getenv("WORKFLOW_ID"),
		RunId:       os.Getenv("RUN_ID"),
		StepName:    os.Getenv("STEP_NAME"),
		StepVersion: os.Getenv("STEP_VERSION"),
	}
}
