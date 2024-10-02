package commit

import (
	"github.com/xanzy/go-gitlab"
	"time"
)

type Commit struct {
	Project    int
	Identifier string
	Title      string
	Message    string
	Link       string
	Author     string
	Date       *time.Time

	Raw *gitlab.Commit
}
