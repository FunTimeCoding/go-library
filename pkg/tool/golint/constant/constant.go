package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"golint",
	"Go source linter with auto-fix support",
	"golint [flags] [path...]",
)
