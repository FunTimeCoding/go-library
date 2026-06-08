package resolve_result

func (r *Result) Identifier() string {
	if len(r.Matches) == 0 {
		return ""
	}

	return r.Matches[0].Identifier
}
