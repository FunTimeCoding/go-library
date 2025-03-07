package merge_request

import "time"

func (r *Request) Age() time.Duration {
	return time.Since(*r.Create)
}
