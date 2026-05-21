package ensure_result

func New(name string) *Result {
	return &Result{Callsign: name, New: true}
}
