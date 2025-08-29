package issue

import (
	"github.com/andygrunwald/go-jira"
	library "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func CommentTime(c *jira.Comment) time.Time {
	return library.PickLatest(ConvertTime(c.Created), ConvertTime(c.Updated))
}
