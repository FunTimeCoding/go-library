package run

import "fmt"

func (r *Run) Environment(
	k string,
	v string,
) {
	r.environment = append(r.environment, fmt.Sprintf("%s=%s", k, v))
}
