package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/inventory"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Store struct {
	inventory *inventory.Inventory
	pools     map[string]*pgxpool.Pool
	sessions  sync.Map
	mu        sync.Mutex
}
