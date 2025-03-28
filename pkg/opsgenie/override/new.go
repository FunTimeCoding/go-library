package override

import "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"

func New(v *schedule.ScheduleOverride) *Override {
	return &Override{
		User:      v.User.Username,
		Start:     v.StartDate,
		End:       v.EndDate,
		Rotations: v.Rotations,
		Raw:       v,
	}
}
