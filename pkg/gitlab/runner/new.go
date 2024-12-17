package runner

import "gitlab.com/gitlab-org/api/client-go"

func New(v *gitlab.Runner) *Runner {
	return &Runner{
		Identifier:  v.ID,
		Name:        v.Name,
		Description: v.Description,
		Address:     v.IPAddress,
		Status:      v.Status,
		Type:        v.RunnerType,
		Raw:         v,
	}
}
