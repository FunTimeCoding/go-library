package alert

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (a *Alert) TeamKeyFallback() string {
	if a.TeamKey == constant.NoKey {
		return a.Team.Name
	}

	return a.TeamKey
}
