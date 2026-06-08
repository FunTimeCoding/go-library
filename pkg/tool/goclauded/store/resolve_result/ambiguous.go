package resolve_result

func (r *Result) Ambiguous() bool {
	return len(r.Matches) > 1
}
