package constant

const TokenEnvironment = "ANTHROPIC_TOKEN" // #nosec G101 not a hardcoded secret

type Mode int

const (
	None Mode = iota
	FiveMinute
	OneHour
)
