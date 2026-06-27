package face

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/inventory"

type Service interface {
	Instances() []inventory.Instance
	Instance(name string) (*inventory.Instance, bool)
	ResolveInstance(explicit string) (string, error)
	ActiveInstance(sessionIdentifier string) (string, bool)
	SetActiveInstance(sessionIdentifier string, instance string)
	Client(instance string) (ProxmoxClient, error)
	SSHClient(instance string) (SnippetClient, error)
}
