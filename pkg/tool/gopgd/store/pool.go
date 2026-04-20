package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *Store) Pool(instance string) (*pgxpool.Pool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p, ok := s.pools[instance]; ok {
		return p, nil
	}

	i, ok := s.Instance(instance)

	if !ok {
		return nil, fmt.Errorf("unknown instance: %s", instance)
	}

	p, e := pgxpool.New(context.Background(), i.Locator())

	if e != nil {
		return nil, fmt.Errorf("connect to %s: %w", instance, e)
	}

	s.pools[instance] = p

	return p, nil
}
