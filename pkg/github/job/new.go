package job

import "github.com/google/go-github/v69/github"

func New(v *github.WorkflowJob) *Job {
	return &Job{Name: v.GetName(), CreatedAt: v.CreatedAt.Time, Raw: v}
}
