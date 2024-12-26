package session_store

import "sync"

type Store struct {
	sync.Mutex
	sessions map[string]string
}
