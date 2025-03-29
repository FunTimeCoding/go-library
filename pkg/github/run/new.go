package run

import "github.com/google/go-github/v70/github"

func New(v *github.WorkflowRun) *Run {
	return &Run{Name: v.GetName(), CreatedAt: v.CreatedAt.Time, Raw: v}
}
