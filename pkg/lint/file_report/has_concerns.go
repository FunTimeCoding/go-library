package file_report

func (r *Report) HasConcerns() bool {
	return len(r.Concerns) > 0
}
