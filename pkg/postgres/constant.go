package postgres

const (
	LocatorEnvironment = "POSTGRES_LOCATOR"

	Command         = "psql"
	UserArgument    = "--username"
	CommandArgument = "--command"
	FileArgument    = "--file"
	EchoAllFlag     = "--echo-all"

	DumpCommand = "pg_dump"
)
