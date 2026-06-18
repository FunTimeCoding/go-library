package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/gosourced/inventory"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	sessions  sync.Map
}
