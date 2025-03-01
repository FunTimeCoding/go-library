package issue

import "k8s.io/apimachinery/pkg/util/duration"

func (i *Issue) formatAge() string {
	return duration.HumanDuration(i.Age())
}
