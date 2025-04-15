package override

import "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"

func NewSlice(v []schedule.ScheduleOverride) []*Override {
	var result []*Override

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
