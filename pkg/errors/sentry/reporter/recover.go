package reporter

func (r *Reporter) Recover(v any) {
	if r.hub == nil {
		return
	}

	r.hub.Recover(v)
}
