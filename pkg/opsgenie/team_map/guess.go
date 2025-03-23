package team_map

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (m *Map) Guess(
	a *alert.Alert,
	verbose bool,
) *team.Team {
	var name string

	if m.tagsToName != nil {
		name = m.tagsToName(a.Tags)
	}

	if name == "" {
		if verbose {
			fmt.Printf("No team: %+v\n", a)
		}

		return nil
	}

	for _, t := range m.Teams {
		if t.Name == name {
			return t
		}
	}

	if verbose {
		fmt.Printf("Guess fail: %+v\n", a)
	}

	return nil
}
