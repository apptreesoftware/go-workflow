package step

import (
	"github.com/apptreesoftware/go-workflow/pkg/core"
	"os"
)

func GetEnvironment() *core.Environment {
	return &core.Environment{
		App:         os.Getenv("PROJECT"),
		WorkflowId:  os.Getenv("WORKFLOW_ID"),
		RunId:       os.Getenv("RUN_ID"),
		StepName:    os.Getenv("STEP_NAME"),
		StepVersion: os.Getenv("STEP_VERSION"),
	}
}
