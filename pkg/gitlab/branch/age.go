package branch

import "time"

func (b *Branch) Age() time.Duration {
	return time.Since(*b.CommitDate)
}
