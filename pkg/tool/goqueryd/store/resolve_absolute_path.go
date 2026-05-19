package store

import "path/filepath"

func (s *Store) ResolveAbsolutePath(
	collection string,
	relativePath string,
) string {
	base, _ := s.collectionPath(collection)

	return filepath.Join(base, relativePath)
}
