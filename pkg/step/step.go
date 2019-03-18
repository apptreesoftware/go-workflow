package step

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/core"
)

type Step interface {
	Name() string
	Version() string
	Execute(in Context) (interface{}, error)
}

func Id(s Step) string {
	return fmt.Sprintf("%s@%s", s.Name(), s.Version())
}

func StepNameAndVersion(environment *core.Environment) string {
	return fmt.Sprintf("%s@%s", environment.StepName, environment.StepVersion)
}

var stepReg = map[string]Step{}

func Register(step Step) {
	id := Id(step)
	stepReg[id] = step
}

func GetStep(id string) Step {
	return stepReg[id]
}
