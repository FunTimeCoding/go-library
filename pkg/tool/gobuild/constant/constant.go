package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gobuild",
	"Cross-compilation and deploy tool",
	"gobuild [flags] [name]",
)
