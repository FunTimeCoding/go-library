package run

import "github.com/google/go-github/v70/github"

func New(v *github.WorkflowRun) *Run {
	return &Run{
		Identifier: v.GetID(),
		Name:       v.GetName(),
		Status:     v.GetStatus(),
		CreatedAt:  v.CreatedAt.Time,
		Raw:        v,
	}
}
