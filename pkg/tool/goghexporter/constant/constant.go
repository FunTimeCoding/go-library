package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goghexporter",
	"GitHub container registry exporter",
	"goghexporter [flags]",
)

const (

	OwnerEnvironment = "GITHUB_EXPORTER_OWNER"
)
