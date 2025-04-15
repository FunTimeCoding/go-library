package issue

import "github.com/andygrunwald/go-jira"

func (i *Issue) LatestComment() *jira.Comment {
	if i.Raw.Fields == nil {
		return nil
	}

	if i.Raw.Fields.Comments == nil {
		return nil
	}

	if c := len(i.Raw.Fields.Comments.Comments); c > 0 {
		return i.Raw.Fields.Comments.Comments[c-1]
	}

	return nil
}
