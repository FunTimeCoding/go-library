package merge_request

func (r *Request) Done() bool {
	if r.State == MergedState {
		return true
	}

	if r.State == ClosedState {
		return true
	}

	return false
}
