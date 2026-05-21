package resolve_result

type Result struct {
	Matches []*Match
}

func (r *Result) Found() bool {
	return len(r.Matches) == 1
}

func (r *Result) Ambiguous() bool {
	return len(r.Matches) > 1
}

func (r *Result) Identifier() string {
	if len(r.Matches) == 0 {
		return ""
	}

	return r.Matches[0].Identifier
}
