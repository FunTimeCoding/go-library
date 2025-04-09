package status

func (s *Status) RawList(a any) *Status {
	return s.Raw(a, "RawList")
}
