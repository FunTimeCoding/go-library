package label_filter

func New(defaultKeep bool) *Filter {
	return &Filter{KeepByDefault: defaultKeep}
}
