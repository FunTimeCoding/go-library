package session

func (s *Session) AliasValue() string {
	if s.Alias == nil {
		return ""
	}

	return *s.Alias
}
