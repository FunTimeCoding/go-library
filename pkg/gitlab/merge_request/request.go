package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"gitlab.com/gitlab-org/api/client-go"
	"time"
)

type Request struct {
	Project    int64
	Identifier int64
	Title      string
	State      string
	Link       string
	Create     *time.Time

	ageColor face.SprintFunction

	Raw *gitlab.BasicMergeRequest
}
