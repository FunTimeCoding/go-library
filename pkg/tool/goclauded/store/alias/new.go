package alias

func New(
	sessionIdentifier string,
	name string,
) *Alias {
	return &Alias{SessionIdentifier: sessionIdentifier, Name: name}
}
