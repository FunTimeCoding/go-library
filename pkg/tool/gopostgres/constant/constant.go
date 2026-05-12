package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gopostgres",
	"PostgreSQL query CLI for gopostgresd",
	"gopostgres [command]",
)

const HostEnvironment = "POSTGRES_HOST"
