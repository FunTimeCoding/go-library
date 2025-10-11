package locator

func (l *Locator) Pointer() *string {
	result := l.String()

	return &result
}
