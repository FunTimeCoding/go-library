package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
)

func (a *Alert) enrich(
	p *option.Alert,
	t *team.Team,
) *Alert {
	a.MonitorIdentifier = fmt.Sprintf(
		"%s-%d",
		constant.OpsgeniePrefix,
		a.SmallIdentifier,
	)
	a.TeamKey = p.Team.KeyByIdentifier(t.Identifier)
	a.Team = t
	a.TeamMap = p.Team
	a.UserMap = p.User
	a.Link = link(a.Identifier, p.WebHost)
	a.shortUser = p.ShortUser
	a.shortAlert = p.ShortAlert
	a.descriptionToName = p.DescriptionToName
	a.parseDescription = p.ParseDescription
	a.Name = a.findName()
	a.Prometheus = a.descriptionToDetail()

	return a
}
