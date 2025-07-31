package event

import (
	"github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	event "k8s.io/api/events/v1"
)

func New(
	v *event.Event,
	cluster string,
) *Event {
	return &Event{
		MonitorIdentifier: constant.GoKevt.StringIdentifier(
			v.Name,
		),
		Cluster:       cluster,
		Namespace:     v.Namespace,
		RegardingKind: v.Regarding.Kind,
		Name:          v.Name,
		Type:          v.Type,
		Reason:        v.Reason,
		Create:        &v.CreationTimestamp.Time,
		Note:          v.Note,
		Raw:           v,
	}
}
