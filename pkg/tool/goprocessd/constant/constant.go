package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goprocessd",
	"Process manager with environment reload",
	"goprocessd [command]",
)

const HistoryCapacity = 200
