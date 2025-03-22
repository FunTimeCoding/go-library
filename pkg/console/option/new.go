package option

func New(v ...Option) *Output {
	result := &Output{Format: Text, Debug: false}

	for _, o := range v {
		o(result)
	}

	return result
}
