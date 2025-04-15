package rotation

import "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"

func NewSlice(v []schedule.Rotation) []*Rotation {
	var result []*Rotation

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
