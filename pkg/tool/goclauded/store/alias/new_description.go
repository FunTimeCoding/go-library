package alias

func NewDescription(
	sessionIdentifier string,
	description string,
) *Alias {
	return &Alias{
		SessionIdentifier: sessionIdentifier,
		Description:       description,
	}
}
