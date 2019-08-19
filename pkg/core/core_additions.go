package core

import (
	"fmt"
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
