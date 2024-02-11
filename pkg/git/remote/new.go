package remote

func New(
	name string,
	locator string,
	provider string,
) *Remote {
	return &Remote{Name: name, Locator: locator, Provider: provider}
}
