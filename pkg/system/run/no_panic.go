package run

func (r *Run) NoPanic() *Run {
	r.Panic = false

	return r
}
