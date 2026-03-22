package run

func (r *Run) WithStdio() *Run {
	r.stdio = true

	return r
}
