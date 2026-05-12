package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gocat",
	"Go file concatenator with import merging",
	"gocat <files...>",
)
