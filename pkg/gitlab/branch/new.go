package branch

import "gitlab.com/gitlab-org/api/client-go"

func New(v *gitlab.Branch) *Branch {
	return &Branch{
		Name: v.Name,
		Raw:  v,
	}
}
