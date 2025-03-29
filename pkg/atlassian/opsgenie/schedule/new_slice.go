package schedule

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team_map"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func NewSlice(
	v []schedule.Schedule,
	m *team_map.Map,
) []*Schedule {
	var result []*Schedule

	for _, element := range v {
		result = append(
			result,
			New(
				&element,
				m.ByIdentifier(element.OwnerTeam.Id),
			),
		)
	}

	return result
}
