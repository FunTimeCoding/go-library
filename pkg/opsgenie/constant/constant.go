package constant

const (
	UserKeyEnvironment = "OPSGENIE_USER_KEY"
	TeamKeyEnvironment = "OPSGENIE_TEAM_KEY"
	WebHostEnvironment = "OPSGENIE_WEB_HOST"

	PageLimit int = 100

	NoKey = "no key"
)

type (
	StringAlias func(s string) string
	SliceAlias  func(a []string) string
)
