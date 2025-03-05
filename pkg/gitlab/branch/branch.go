package branch

import "gitlab.com/gitlab-org/api/client-go"

type Branch struct {
	Name string

	Raw *gitlab.Branch
}
