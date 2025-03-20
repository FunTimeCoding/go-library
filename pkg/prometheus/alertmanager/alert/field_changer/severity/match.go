package severity

func (s *Severity) Match(
	alert string,
	name string,
) bool {
	return s.Alert == alert && s.From == name
}
