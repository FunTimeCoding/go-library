package event

import (
	event "k8s.io/api/events/v1"
	"time"
)

type Event struct {
	MonitorIdentifier string
	Cluster           string
	Namespace         string
	RegardingKind     string
	Name              string
	Type              string
	Reason            string
	Note              string
	Create            *time.Time

	Raw *event.Event
}
