package event

import "k8s.io/api/events/v1"

func NewSlice(
	v []v1.Event,
	cluster string,
) []*Event {
	var result []*Event

	for _, e := range v {
		result = append(result, New(&e, cluster))
	}

	return result
}
