package session

func (s *Session) CallsignValue() string {
	if s.Callsign == nil {
		return ""
	}

	return *s.Callsign
}
