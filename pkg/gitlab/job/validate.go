package job

import (
	"slices"
	"strings"
)

func (j *Job) Validate() {
	if j.Fail() && !slices.Contains(j.concern, FailConcern) {
		j.concern = append(j.concern, FailConcern)
	}

	if j.Trace != "" && !slices.Contains(j.concern, Timeout) {
		if strings.Contains(j.Trace, traceTimeoutMatch) {
			j.concern = append(j.concern, Timeout)
		}
	}
}
