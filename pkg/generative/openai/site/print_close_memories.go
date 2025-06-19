package site

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium/protocol"
)

func (s *Site) printCloseMemories() {
	s.protocol.PrintNode(CloseDialogSelector, usefulAttributes)
	n := s.protocol.Select(CloseDialogSelector, 2)
	fmt.Println("Close dialog index 2")
	protocol.Print(n, usefulAttributes)
}
