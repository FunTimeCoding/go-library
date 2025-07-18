package job

import "time"

func (j *Job) Age() time.Duration {
	return time.Since(*j.Create)
}
