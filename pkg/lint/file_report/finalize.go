package file_report

func (r *Report) Finalize() *Report {
	r.Original = r.originalBuilder.String()

	if r.Modified() {
		r.Fixed = r.fixedBuilder.String()
	}

	return r
}
