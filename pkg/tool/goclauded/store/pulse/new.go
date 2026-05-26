package pulse

func New(
	sessionIdentifier string,
	fromName string,
	body string,
) *Pulse {
	return &Pulse{
		SessionIdentifier: sessionIdentifier,
		FromName:          fromName,
		Body:              body,
	}
}
