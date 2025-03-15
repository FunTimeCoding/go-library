package silence

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (s *Silence) Expired() bool {
	return s.State == constant.ExpiredState
}
