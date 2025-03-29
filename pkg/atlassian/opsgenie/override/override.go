package override

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"time"
)

type Override struct {
	User      string
	Start     time.Time
	End       time.Time
	Rotations []schedule.RotationIdentifier

	Raw *schedule.ScheduleOverride
}
