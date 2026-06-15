package event_metadata

func New(
	eventIdentifier uint,
	key string,
	value string,
) *EventMetadata {
	return &EventMetadata{
		EventIdentifier: eventIdentifier,
		Key:             key,
		Value:           value,
	}
}
