package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"gitlab.com/gitlab-org/api/client-go"
	"time"
)

type Request struct {
	Project    int
	Identifier int
	Title      string
	State      string
	Link       string
	Create     *time.Time

	AgeColor face.SprintFunction

	Raw *gitlab.BasicMergeRequest
}
