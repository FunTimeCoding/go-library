package job

import "github.com/funtimecoding/go-library/pkg/gitlab/constant"

func (j *Job) Fail() bool {
	return j.Status == constant.Failed
}
