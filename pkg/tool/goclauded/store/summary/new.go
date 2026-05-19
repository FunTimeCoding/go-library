package summary

func New(
	sessionIdentifier string,
	name string,
	body string,
) *Summary {
	return &Summary{
		SessionIdentifier: sessionIdentifier,
		Name:              name,
		Body:              body,
	}
}
