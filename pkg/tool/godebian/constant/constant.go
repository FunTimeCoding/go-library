package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"godebian",
	"Debian package builder",
	"godebian <executable> <version>",
)
