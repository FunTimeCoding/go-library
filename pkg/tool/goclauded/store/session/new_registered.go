package session

import "time"

func NewRegistered(
	identifier string,
	callsign string,
) *Session {
	now := time.Now()

	return &Session{
		Identifier: identifier,
		Name:       callsign,
		Callsign:   &callsign,
		LastSeen:   now,
		StartedAt:  now,
	}
}
