package silence

import "time"

func (s *Silence) Remain() time.Duration {
	return time.Until(*s.End)
}
