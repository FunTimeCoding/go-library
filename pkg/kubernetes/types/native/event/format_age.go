package event

import "k8s.io/apimachinery/pkg/util/duration"

func (e *Event) formatAge() string {
	return duration.HumanDuration(e.Age())
}
