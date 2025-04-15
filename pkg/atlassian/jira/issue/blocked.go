package issue

func (i *Issue) Blocked() bool {
	for _, element := range i.Raw.Fields.IssueLinks {
		if element.InwardIssue == nil {
			continue
		}

		if element.Type.Inward == BlockedBy {
			return true
		}
	}

	return false
}
