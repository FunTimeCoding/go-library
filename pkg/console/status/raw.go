package status

func (s *Status) Raw(a any) *Status {
	return s.Line("  Raw: %+v", a)
}
