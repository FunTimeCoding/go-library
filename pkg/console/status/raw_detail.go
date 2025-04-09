package status

func (s *Status) RawDetail(a any) *Status {
	return s.Raw(a, "RawDetail")
}
