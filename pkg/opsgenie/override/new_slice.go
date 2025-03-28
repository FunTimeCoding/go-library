package override

import "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"

func NewSlice(v []schedule.ScheduleOverride) []*Override {
	var result []*Override

	for _, element := range v {
		result = append(result, New(&element))
	}

	return result
}
