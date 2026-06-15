package timeline

func (e *Entry) meta(key string) string {
	if e.Metadata == nil {
		return ""
	}

	return e.Metadata[key]
}
