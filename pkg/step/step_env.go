package step

import (
	"github.com/apptreesoftware/go-workflow/pkg/core"
	"os"
)

func GetIPCEnvironment() *core.Environment {
	return &core.Environment{
		Project:     os.Getenv("PROJECT"),
		Workflow:    os.Getenv("WORKFLOW_ID"),
		RunId:       os.Getenv("RUN_ID"),
		StepName:    os.Getenv("STEP_NAME"),
		StepVersion: os.Getenv("STEP_VERSION"),
		CacheHost:   os.Getenv("WORKFLOW_CACHE_CONNECTION"),
		Debug:       os.Getenv("DEBUG") == "true",
	}
}
