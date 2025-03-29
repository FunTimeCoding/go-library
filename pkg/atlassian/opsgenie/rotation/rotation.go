package rotation

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"time"
)

type Rotation struct {
	Identifier   string
	Name         string
	Type         og.RotationType
	Participants []og.Participant
	Start        *time.Time

	Raw *schedule.Rotation
}
