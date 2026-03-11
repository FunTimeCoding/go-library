package file_report

func (r *Report) PassLine(s string) {
	r.append(r.fixedBuilder, s)
}
