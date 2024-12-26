package session_store

func New() *Store {
	return &Store{sessions: make(map[string]string)}
}
