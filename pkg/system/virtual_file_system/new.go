package virtual_file_system

func New() *System {
	return &System{files: make(map[string]string)}
}
