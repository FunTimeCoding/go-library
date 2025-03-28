package rotation

import "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"

func NewSlice(v []schedule.Rotation) []*Rotation {
	var result []*Rotation

	for _, element := range v {
		result = append(result, New(&element))
	}

	return result
}
