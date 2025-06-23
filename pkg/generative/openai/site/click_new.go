package site

import "time"

func (s *Site) clickNew() {
	n := s.protocol.Select(NewSelector, 0)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
	time.Sleep(1 * time.Second)
}
