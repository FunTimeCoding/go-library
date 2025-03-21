package rule_list

func (l *List) Runbook(name string) string {
	if r := l.Find(name); r != nil {
		return r.Runbook
	}

	return ""
}
