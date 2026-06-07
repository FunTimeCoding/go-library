package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"

func (s *Store) SendMessage(
	fromName string,
	toName string,
	body string,
) error {
	return s.database.Create(message.New(fromName, toName, body)).Error
}
