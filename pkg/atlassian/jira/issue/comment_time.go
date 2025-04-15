package issue

import (
	"github.com/andygrunwald/go-jira"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func CommentTime(c *jira.Comment) time.Time {
	return timeLibrary.PickLatest(
		ConvertTime(c.Created),
		ConvertTime(c.Updated),
	)
}
