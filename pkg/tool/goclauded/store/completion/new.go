package completion

func New(
	sessionIdentifier string,
	name string,
	kind string,
	topic string,
	summary string,
) *Completion {
	return &Completion{
		SessionIdentifier: sessionIdentifier,
		Name:              name,
		Kind:              kind,
		Topic:             topic,
		Summary:           summary,
	}
}
