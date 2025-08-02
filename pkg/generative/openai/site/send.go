package site

import "fmt"

func (s *Site) Send(t string) {
	if false {
		fmt.Printf("Focused: %+v\n", s.protocol.Focused())

		return
	}

	s.protocol.EnterText(`#prompt-textarea`, t)
}
