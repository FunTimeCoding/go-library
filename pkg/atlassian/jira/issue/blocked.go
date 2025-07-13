package issue

import "slices"

func (i *Issue) Blocked() bool {
	for _, l := range i.Raw.Fields.IssueLinks {
		if l.InwardIssue == nil {
			continue
		}

		if l.Type.Inward == BlockedBy {
			if slices.Contains(
				i.option.DoneStatus,
				l.InwardIssue.Fields.Status.Name,
			) {
				continue
			}

			return true
		}
	}

	return false
}
