package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goterraformd",
	"Terraform apply runner",
	"goterraformd",
).WithInstructions(
	"Terraform apply runner - trigger init + apply and check results. Runs execute asynchronously; use runs and run tools to poll for completion.",
)

const (
	Trigger    = "trigger"
	Runs       = "runs"
	Run        = "run"
	Limit      = "limit"
	Status     = "status"
	Identifier = "id"

	RecentRunsFailed = "failed to list recent runs"
	RunLookupFailed  = "failed to look up run"
)
