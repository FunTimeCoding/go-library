package site

import "fmt"

func (s *Site) Dump() {
	fmt.Println(s.protocol.Outer("body"))
}
