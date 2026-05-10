package constant

const (
	HostEnvironment           = "SALT_HOST"
	PortEnvironment           = "SALT_PORT"
	UserEnvironment           = "SALT_USER"
	PasswordEnvironment       = "SALT_PASSWORD"
	AuthenticationEnvironment = "SALT_AUTHENTICATION"

	Run       = "cmd.run"
	Highstate = "state.highstate"

	TokenHeader = "X-Auth-Token"

	LoginPath   = "login"
	KeysPath    = "keys"
	MinionsPath = "minions"
	JobsPath    = "jobs"

	LocalClient      = "local"
	LocalAsyncClient = "local_async"
	WheelClient      = "wheel"

	GlobTarget = "glob"

	KeyAccept = "key.accept"
	KeyDelete = "key.delete"
)
