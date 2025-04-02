package loader

func New() *Loader {
	return &Loader{contents: make(map[string]string)}
}
