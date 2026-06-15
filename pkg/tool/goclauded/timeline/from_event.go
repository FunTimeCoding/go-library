package timeline

import "time"

func FromEvent(
	identifier uint,
	sessionIdentifier string,
	kind string,
	actor string,
	metadata map[string]string,
	createdAt time.Time,
) *Entry {
	return &Entry{
		Identifier:        identifier,
		SessionIdentifier: sessionIdentifier,
		Timestamp:         createdAt.UTC().Format(time.RFC3339),
		Kind:              kind,
		Actor:             actor,
		Metadata:          metadata,
	}
}
