package output

func FromOptions(o ...Option) *Settings {
	s := Ensure(nil)

	for _, element := range o {
		element(s)
	}

	return s
}
