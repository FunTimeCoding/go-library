package resolve_result

func (r *Result) Found() bool {
	return len(r.Matches) == 1
}
