package run

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github/repository"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/google/go-github/v70/github"
	"time"
)

func New(v *github.WorkflowRun) *Run {
	var create time.Time
	var update time.Time

	if v.CreatedAt != nil {
		create = v.CreatedAt.Time
	}

	if v.UpdatedAt == nil {
		update = create
	} else {
		update = v.UpdatedAt.Time
	}

	return &Run{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%d",
			constant.GitHubPrefix,
			v.GetID(),
		),
		Identifier: v.GetID(),
		Name:       v.GetName(),
		Status:     v.GetStatus(),
		Create:     create,
		Update:     update,
		Repository: repository.New(v.Repository),
		Raw:        v,
	}
}
