package workflow

import "github.com/google/go-github/v81/github"

func New(v *github.Workflow) *Workflow {
	return &Workflow{
		Identifier: v.GetID(),
		Name:       v.GetName(),
		State:      v.GetState(),
		CreatedAt:  v.CreatedAt.Time,
		Raw:        v,
	}
}
