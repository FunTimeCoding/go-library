package output

func Ensure(s *Settings) *Settings {
	if s == nil {
		return Default
	}

	return s
}
