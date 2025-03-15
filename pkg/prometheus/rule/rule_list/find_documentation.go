package rule_list

func (l *List) FindDocumentation(name string) string {
	if r := l.Find(name); r != nil {
		return r.Documentation
	}

	return ""
}
