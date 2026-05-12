package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gosaltd",
	"Salt highstate runner",
	"gosaltd",
).WithInstructions(
	"Salt highstate runner - trigger highstate on all minions and check results. Runs execute asynchronously; use runs and run tools to poll for completion.",
)

const (
	Trigger    = "trigger"
	Runs       = "runs"
	Run        = "run"
	Minions    = "minions"
	Keys       = "keys"
	Limit      = "limit"
	Status     = "status"
	Identifier = "id"

	RecentRunsFailed = "failed to list recent runs"
	RunLookupFailed  = "failed to look up run"
	MinionsFailed    = "failed to list minions"
	KeysFailed       = "failed to list keys"
	NotConnected     = "salt API not connected yet"
)
