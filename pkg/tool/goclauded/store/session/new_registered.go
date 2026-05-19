package session

import "time"

func NewRegistered(
	identifier string,
	name string,
) *Session {
	now := time.Now()

	return &Session{
		Identifier:  identifier,
		Name:        name,
		NeedsRoster: true,
		LastSeen:    now,
		CreatedAt:   now,
	}
}
