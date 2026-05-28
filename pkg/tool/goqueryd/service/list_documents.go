package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/result"

func (s *Service) ListDocuments(
	collection string,
) ([]result.DocumentEntry, error) {
	return s.store.ListDocuments(collection)
}
