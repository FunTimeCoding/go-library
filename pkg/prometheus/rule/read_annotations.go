package rule

func (a *Rule) readAnnotations() {
	if a.RawAlert == nil {
		return
	}

	for k, v := range a.RawAlert.Annotations {
		switch k {
		case DocumentationKey:
			a.Documentation = string(v)
		case RunbookLocatorKey:
			a.Documentation = string(v)
		}

		if a.Documentation != "" {
			break
		}
	}
}
