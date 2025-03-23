package alert

import "k8s.io/apimachinery/pkg/util/duration"

func (a *Alert) formatAge() string {
	// Even denser than units.HumanDuration
	return duration.HumanDuration(a.Age())
}
