package filter

func New() *Filter {
	return &Filter{
		equal: make(map[string]string),
		like:  make(map[string]string),
	}
}
