package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/notification"

func (s *Store) SendNotification(
	callsign string,
	source string,
	body string,
) error {
	return s.database.Create(notification.New(callsign, source, body)).Error
}
