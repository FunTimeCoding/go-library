package constant

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/detail"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

const (
	UserKeyEnvironment = "OPSGENIE_USER_KEY"
	TeamKeyEnvironment = "OPSGENIE_TEAM_KEY"
	WebHostEnvironment = "OPSGENIE_WEB_HOST"

	PageLimit int = 100

	NoKey = "no key"
)

var Format = option.ExtendedColor.Copy().Tag(
	tag.Link,
	tag.Category,
)

type ParseDescription func(string) *detail.Prometheus
