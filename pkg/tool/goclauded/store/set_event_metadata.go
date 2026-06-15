package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event_metadata"

func (s *Store) SetEventMetadata(
	eventIdentifier uint,
	metadata map[string]string,
) error {
	for key, value := range metadata {
		if e := s.database.Create(
			event_metadata.New(eventIdentifier, key, value),
		).Error; e != nil {
			return e
		}
	}

	return nil
}
