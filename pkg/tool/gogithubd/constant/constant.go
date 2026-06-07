package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gogithubd",
	"GitHub container registry exporter",
	"gogithubd [flags]",
)

const (
	OwnerEnvironment = "GITHUB_EXPORTER_OWNER"
)
