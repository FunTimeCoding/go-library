package label

func New(
	sessionIdentifier string,
	key string,
	value string,
) *Label {
	return &Label{
		SessionIdentifier: sessionIdentifier,
		Key:               key,
		Value:             value,
	}
}
