package job

import "github.com/google/go-github/v70/github"

func New(v *github.WorkflowJob) *Job {
	return &Job{
		Identifier: v.GetID(),
		Name:       v.GetName(),
		Status:     v.GetStatus(),
		Hash:       v.GetHeadSHA(),
		CreatedAt:  v.CreatedAt.Time,
		Raw:        v,
	}
}
