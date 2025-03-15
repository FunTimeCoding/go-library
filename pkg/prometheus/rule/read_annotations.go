package rule

func (r *Rule) readAnnotations() {
	if r.RawAlert == nil {
		return
	}

	for k, v := range r.RawAlert.Annotations {
		switch k {
		case DocumentationKey:
			r.Documentation = string(v)
		case RunbookLocatorKey:
			r.Documentation = string(v)
		}

		if r.Documentation != "" {
			break
		}
	}
}
