package file_report

func (r *Report) ProcessConcerns(fix bool) {
	if r.HasConcerns() {
		if fix && r.Fix != nil {
			r.Fix()
		}
	}
}
