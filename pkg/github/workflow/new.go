package workflow

import "github.com/google/go-github/v70/github"

func New(v *github.Workflow) *Workflow {
	return &Workflow{
		Identifier: v.GetID(),
		Name:       v.GetName(),
		CreatedAt:  v.CreatedAt.Time,
		Raw:        v,
	}
}
