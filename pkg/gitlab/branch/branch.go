package branch

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"gitlab.com/gitlab-org/api/client-go"
	"time"
)

type Branch struct {
	Name       string
	Merged     bool
	CommitDate *time.Time

	AgeColor face.SprintFunction

	Raw *gitlab.Branch
}
