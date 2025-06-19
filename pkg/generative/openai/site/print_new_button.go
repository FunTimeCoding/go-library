package site

import "github.com/funtimecoding/go-library/pkg/chromium/protocol"

func (s *Site) printNewButton() {
	s.protocol.PrintNode(NewChatSelector, usefulAttributes)
	protocol.Print(
		s.protocol.Select(NewChatSelector, 0),
		usefulAttributes,
	)
}
