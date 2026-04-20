package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Store struct {
	configuration *config.Configuration
	pools         map[string]*pgxpool.Pool
	sessions      sync.Map
	mu            sync.Mutex
}
