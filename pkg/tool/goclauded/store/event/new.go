package event

func New(
	sessionIdentifier string,
	kind string,
	actor string,
) *Event {
	return &Event{
		SessionIdentifier: sessionIdentifier,
		Kind:              kind,
		Actor:             actor,
	}
}
