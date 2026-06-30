package constant

const TokenEnvironment = "ANTHROPIC_TOKEN" // #nosec G101 not a hardcoded secret
const VocabularyLink = "https://raw.githubusercontent.com/rohangpta/ctoc/main/vocab.json"

type Mode int

const (
	None Mode = iota
	FiveMinute
	OneHour
)
