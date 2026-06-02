package service

func (s *Service) Sync(files []DiscoveredFile) {
	var paths []string

	for _, f := range files {
		paths = append(paths, f.Path)
		s.store.UpsertSeed(f.Name, f.Path, f.ContentHash, f.Content)
	}

	s.store.RemoveMissing(paths)
	s.notifier.Notify()
}
