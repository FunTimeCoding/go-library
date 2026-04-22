package option

import "github.com/funtimecoding/go-library/pkg/tool/gopgd/inventory"

type Postgres struct {
	Port      int
	Inventory *inventory.Inventory
}
