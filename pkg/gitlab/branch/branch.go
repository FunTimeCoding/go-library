package branch

import (
	"gitlab.com/gitlab-org/api/client-go"
	"time"
)

type Branch struct {
	Name       string
	Merged     bool
	CommitDate *time.Time

	Raw *gitlab.Branch
}
