package run

import "github.com/google/go-github/v69/github"

func New(v *github.WorkflowRun) *Run {
	return &Run{Name: v.GetName(), CreatedAt: v.CreatedAt.Time, Raw: v}
}
