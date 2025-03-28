package schedule

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/team"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

type Schedule struct {
	Identifier  string
	Name        string
	Description string

	Team *team.Team

	Raw *schedule.Schedule
}
