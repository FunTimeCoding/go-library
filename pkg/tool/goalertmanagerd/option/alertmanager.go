package option

import "github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/inventory"

type Alertmanager struct {
	Port      int
	Version   string
	Inventory *inventory.Inventory
}
