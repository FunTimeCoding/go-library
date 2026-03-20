package issue

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	library "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func (i *Issue) ChangeTime() time.Time {
	var result time.Time

	if u := i.LatestComment(); u != nil {
		result = CommentTime(u)
	} else {
		result = library.PickLatest(
			time.Time(i.Raw.Fields.Created),
			time.Time(i.Raw.Fields.Updated),
		)
	}

	if t := i.lastStatusTransition(); t.After(result) {
		result = t
	}

	return result
}

func (i *Issue) lastStatusTransition() time.Time {
	var result time.Time

	if i.Raw.Changelog == nil {
		return result
	}

	for _, h := range i.Raw.Changelog.Histories {
		for _, item := range h.Items {
			if item.Field == "status" {
				t, e := time.Parse(
					constant.TimeFormat,
					h.Created,
				)

				if e != nil {
					continue
				}

				if t.After(result) {
					result = t
				}
			}
		}
	}

	return result
}
