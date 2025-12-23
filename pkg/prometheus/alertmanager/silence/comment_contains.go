package silence

import "strings"

func (s *Silence) CommentContains(v string) bool {
	return strings.Contains(s.Comment, v)
}
