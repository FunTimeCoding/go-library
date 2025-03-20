package field_changer

func New() *Changer {
	return &Changer{name: make(map[string]string)}
}
