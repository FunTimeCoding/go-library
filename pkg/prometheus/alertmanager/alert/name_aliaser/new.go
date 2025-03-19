package name_aliaser

func New() *Aliaser {
	return &Aliaser{aliases: make(map[string]string)}
}
