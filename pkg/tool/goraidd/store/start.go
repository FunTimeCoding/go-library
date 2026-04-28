package store

func (s *Store) Start() {
	go s.run()
}
