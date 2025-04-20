package issue

func (i *Issue) Blocked() bool {
	for _, l := range i.Raw.Fields.IssueLinks {
		if l.InwardIssue == nil {
			continue
		}

		if l.Type.Inward == BlockedBy {
			return true
		}
	}

	return false
}
