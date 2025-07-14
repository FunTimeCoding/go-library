package job

import "slices"

func (j *Job) HasConcern(s string) bool {
	return slices.Contains(j.concern, s)
}
