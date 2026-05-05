package option

import "github.com/funtimecoding/go-library/pkg/tool/gopostgresd/inventory"

type Postgres struct {
	Port      int
	Inventory *inventory.Inventory
	Version   string
}
