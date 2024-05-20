package sentence

func (s *Sentence) Add(v string) {
	s.affect = append(s.affect, v)
}
