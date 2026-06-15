package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event_metadata"
)

func (s *Store) UpdateEventMetadata(
	eventIdentifier uint,
	key string,
	value string,
) error {
	return s.database.Model(event_metadata.Stub()).Where(
		"event_identifier = ? AND key = ?",
		eventIdentifier,
		key,
	).Update(constant.Value, value).Error
}
