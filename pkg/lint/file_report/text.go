package file_report

func (r *Report) Text() (string, int) {
	r.lineNumber++
	result := r.scanner.Text()
	r.append(r.originalBuilder, result)

	return result, r.lineNumber
}
