package event

func New(
	sessionIdentifier string,
	kind string,
	name string,
	target string,
	body string,
) *Event {
	return &Event{
		SessionIdentifier: sessionIdentifier,
		Kind:              kind,
		Name:              name,
		Target:            target,
		Body:              body,
	}
}
