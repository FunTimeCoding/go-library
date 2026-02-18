package lifecycle

func New(v ...Option) *Lifecycle {
	result := &Lifecycle{}

	for _, f := range v {
		f(result)
	}

	return result
}
