package job

import (
	"github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"gitlab.com/gitlab-org/api/client-go"
)

func New(v *gitlab.Job) *Job {
	return &Job{
		MonitorIdentifier: constant.GoGitLab.IntegerIdentifier(v.ID),
		Identifier:        v.ID,
		Name:              v.Name,
		Status:            v.Status,
		Stage:             v.Stage,
		Create:            v.CreatedAt,
		Link:              v.WebURL,
		Raw:               v,
	}
}
