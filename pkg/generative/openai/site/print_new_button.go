package site

import "github.com/funtimecoding/go-library/pkg/chromium/protocol"

func (s *Site) printNewButton() {
	s.protocol.PrintNode(NewSelector, usefulAttributes)
	protocol.Print(
		s.protocol.Select(NewSelector, 0),
		usefulAttributes,
	)
}
