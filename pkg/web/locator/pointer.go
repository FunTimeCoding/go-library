package locator

func (l *Locator) Pointer() *string {
	return new(l.String())
}
