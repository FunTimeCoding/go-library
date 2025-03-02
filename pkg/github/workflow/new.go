package workflow

import "github.com/google/go-github/v69/github"

func New(v *github.Workflow) *Workflow {
	return &Workflow{Name: v.GetName(), CreatedAt: v.CreatedAt.Time, Raw: v}
}
