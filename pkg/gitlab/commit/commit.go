package commit

import (
	"gitlab.com/gitlab-org/api/client-go"
	"time"
)

type Commit struct {
	Identifier string
	Title      string
	Message    string
	Link       string
	Author     string
	Date       *time.Time

	Raw *gitlab.Commit
}
