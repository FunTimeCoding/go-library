package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gorunif",
	"Conditional command runner",
	"gorunif [flags] <command>",
)
