package alert

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/alert/option"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team"
)

func (a *Alert) enrich(
	p *option.Alert,
	t *team.Team,
) *Alert {
	a.TeamKey = p.Team.KeyByIdentifier(t.Identifier)
	a.Team = t
	a.TeamMap = p.Team
	a.UserMap = p.User
	a.Link = link(a.Identifier, p.WebHost)
	a.shortUser = p.ShortUser
	a.shortAlert = p.ShortAlert
	a.descriptionToName = p.DescriptionToName
	a.Name = a.findName()

	return a
}
