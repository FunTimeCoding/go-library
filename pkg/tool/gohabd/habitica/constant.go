package habitica

const (
	HostEnvironment  = "HABITICA_HOST"
	UserEnvironment  = "HABITICA_USER"
	TokenEnvironment = "HABITICA_TOKEN" // #nosec G101 not a hardcoded secret
)

const (
	apiBase = "/api/v3"

	userHeader  = "x-api-user"
	tokenHeader = "x-api-key"
)
