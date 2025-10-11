package locator

func New(h string) *Locator {
	result := Stub()
	result.host = h

	return result
}
