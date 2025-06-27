package run

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/google/go-github/v70/github"
)

func New(v *github.WorkflowRun) *Run {
	return &Run{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%d",
			constant.GitHubPrefix,
			v.GetID(),
		),
		Identifier: v.GetID(),
		Name:       v.GetName(),
		Status:     v.GetStatus(),
		Create:     v.CreatedAt.Time,
		Update:     v.UpdatedAt.Time,
		Raw:        v,
	}
}
