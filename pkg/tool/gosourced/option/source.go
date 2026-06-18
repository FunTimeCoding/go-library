package option

import "github.com/funtimecoding/go-library/pkg/tool/gosourced/inventory"

type Source struct {
	Port      int
	Version   string
	Inventory *inventory.Inventory
}
