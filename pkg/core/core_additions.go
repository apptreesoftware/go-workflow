package core

import "time"

func (h WorkflowHistory) GetStartTime() time.Time {
	return time.Unix(h.Start, 0)
}

func (h WorkflowHistory) GetEndTime() time.Time {
	return time.Unix(h.End, 0)
}