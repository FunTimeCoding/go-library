package format

func (s *Settings) Tag(v string) *Settings {
	s.Tags = append(s.Tags, v)

	return s
}
