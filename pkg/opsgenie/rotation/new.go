package rotation

import "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"

func New(v *schedule.Rotation) *Rotation {
	return &Rotation{
		Identifier:   v.Id,
		Name:         v.Name,
		Type:         v.Type,
		Start:        v.StartDate,
		Participants: v.Participants,
		Raw:          v,
	}
}
