package output

func New(o ...Option) *Settings {
	result := &Settings{Format: Text, Debug: false}

	for _, element := range o {
		element(result)
	}

	return result
}
