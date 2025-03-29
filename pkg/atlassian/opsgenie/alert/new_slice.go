package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/option"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func NewSlice(
	v []alert.Alert,
	p *option.Alert,
	verbose bool,
) []*Alert {
	var result []*Alert

	for _, a := range v {
		t := p.Team.ByIdentifier(a.OwnerTeamID)

		if t == nil && verbose {
			if t = p.Team.Guess(&a, true); t == nil {
				fmt.Printf("Team not found: %+v\n", a)

				continue
			}
		}

		result = append(result, New(&a, p, t))
	}

	return result
}
