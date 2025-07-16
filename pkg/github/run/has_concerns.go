package run

func (r *Run) HasConcerns() bool {
	return len(r.concern) > 0
}
