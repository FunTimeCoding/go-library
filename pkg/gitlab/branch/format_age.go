package branch

import "k8s.io/apimachinery/pkg/util/duration"

func (b *Branch) formatAge() string {
	return duration.HumanDuration(b.Age())
}
