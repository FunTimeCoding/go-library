package file_report

func (r *Report) ChangedLine(s string) {
	r.append(r.fixedBuilder, s)

	if !r.modified {
		r.modified = true
	}
}
