package virtual_file_system

type System struct {
	files    map[string]*File
	moves    []PendingMove
	written  map[string]bool
	deleted  map[string]bool
	maxFiles int
	maxBytes int64
}
