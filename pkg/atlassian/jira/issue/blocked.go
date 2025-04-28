package issue

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"

func (i *Issue) Blocked() bool {
	for _, l := range i.Raw.Fields.IssueLinks {
		if l.InwardIssue == nil {
			continue
		}

		if l.Type.Inward == BlockedBy {
			if l.InwardIssue.Fields.Status.Name == constant.Closed {
				continue
			}

			return true
		}
	}

	return false
}
