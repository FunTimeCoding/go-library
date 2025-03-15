package silence

import "time"

func (s *Silence) Age() time.Duration {
	return time.Since(*s.Start)
}
