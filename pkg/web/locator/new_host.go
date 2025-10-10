package locator

func NewHost(h string) *Locator {
	result := New()
	result.host = h

	return result
}
