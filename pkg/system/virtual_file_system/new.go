package virtual_file_system

func New(options ...func(*System)) *System {
	result := &System{
		files:   make(map[string]*File),
		written: make(map[string]bool),
		deleted: make(map[string]bool),
	}

	for _, o := range options {
		o(result)
	}

	return result
}
