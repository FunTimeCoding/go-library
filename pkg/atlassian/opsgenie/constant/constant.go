package constant

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/detail"

const (
	UserKeyEnvironment = "OPSGENIE_USER_KEY"
	TeamKeyEnvironment = "OPSGENIE_TEAM_KEY"
	WebHostEnvironment = "OPSGENIE_WEB_HOST"

	PageLimit int = 100

	NoKey = "no key"
)

type ParseDescription func(string) *detail.Prometheus
