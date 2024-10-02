package format

func (s *Settings) Copy() *Settings {
	other := *s

	return &other
}
