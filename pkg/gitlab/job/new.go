package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"gitlab.com/gitlab-org/api/client-go"
)

func New(v *gitlab.Job) *Job {
	return &Job{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%d",
			constant.GitLabPrefix,
			v.ID,
		),
		Identifier: v.ID,
		Name:       v.Name,
		Status:     v.Status,
		Stage:      v.Stage,
		Create:     v.CreatedAt,
		Link:       v.WebURL,
		Raw:        v,
	}
}
