package ensure_result

func New(name string) *Result {
	return &Result{Name: name, New: true}
}
