package file_report

func (r *Report) Scan() bool {
	return r.scanner.Scan()
}
