package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(c *config.Configuration) *Store {
	return &Store{
		configuration: c,
		pools:         make(map[string]*pgxpool.Pool),
	}
}
