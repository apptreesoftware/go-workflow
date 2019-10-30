package core

import (
	"fmt"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

func (m JobSearchResult) GetStartTime() time.Time {
	return time.Unix(m.Start, 0)
}

func (m JobSearchResult) GetEndTime() time.Time {
	return time.Unix(m.End, 0)
}

func (m JobSearchResult) DurationString() string {
	return fmt.Sprintf("%.2f sec", m.Duration)
}

func (m JobSearchResult) SpawnTrailString() (workflowPath string, jobPath string) {
	workflowPathBuilder := strings.Builder{}
	jobPathBuilder := strings.Builder{}
	count := len(m.SpawnTrail)
	for i, entry := range m.SpawnTrail {
		workflowPathBuilder.WriteString(entry.Workflow)
		jobPathBuilder.WriteString(entry.JobId)
		if i < count-1 {
			workflowPathBuilder.WriteString("->")
			jobPathBuilder.WriteString("->")
		}
	}
	return workflowPathBuilder.String(), jobPathBuilder.String()
}

func (m JobFailureInfo) FailureInfoString() string {
	return fmt.Sprintf("Step %d - %s failed. Reason: %s", m.StepIndex, m.StepName, m.Message)
}

func (p Package) IsHostedPackage() bool {
	for _, step := range p.Steps {
		if len(step.Url) > 0 {
			return true
		}
	}
	return false
}

func (p Package) ValidatePackageType() error {
	if p.IsHostedPackage() && p.IsExecutablePackage() {
		return xerrors.Errorf("Package contains both executables and hosted steps. This is not currently supported")
	}
	if p.IsHostedPackage() {
		for name, step := range p.Steps {
			if len(step.Url) == 0 {
				return xerrors.Errorf("The step %s does not contain a `url`. " +
						"All steps must contain a url when publishing a hosted package", name)
			}
		}
	}
	return nil
}


func (p Package) IsExecutablePackage() bool {
	return p.Executables != nil
}