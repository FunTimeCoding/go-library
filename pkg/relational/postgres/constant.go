package postgres

const (
	AdministratorLocatorEnvironment = "POSTGRES_ADMINISTRATOR_LOCATOR"

	LocatorEnvironment = "POSTGRES_LOCATOR"

	Command         = "psql"
	UserArgument    = "--username"
	CommandArgument = "--command"
	FileArgument    = "--file"
	EchoAllFlag     = "--echo-all"

	DumpCommand = "pg_dump"
)
