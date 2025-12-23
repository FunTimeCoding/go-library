package silence

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (s *Silence) IsPermanent() bool {
	return s.CommentContains(constant.PermanentTag)
}
