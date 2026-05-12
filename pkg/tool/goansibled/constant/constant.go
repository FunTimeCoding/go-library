package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"goansibled",
	"Ansible playbook runner",
	"goansibled",
).WithInstructions(
	"Ansible playbook runner - trigger runs and check results. Runs execute asynchronously; use runs and run tools to poll for completion and read output.",
)

const (

	Playbooks          = "playbooks"
	Trigger            = "trigger"
	Runs               = "runs"
	Run                = "run"
	Playbook           = "playbook"
	Limit              = "limit"
	Status             = "status"
	Identifier         = "id"

	RecentRunsFailed = "failed to list recent runs"
	RunLookupFailed  = "failed to look up run"
)
