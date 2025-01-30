package format

func (s *Settings) RemoveTag(v string) {
	for i, tag := range s.Tags {
		if tag == v {
			s.Tags = append(s.Tags[:i], s.Tags[i+1:]...)

			return
		}
	}
}
