package argument

func (i *Instance) PositionalFallback(
	number int,
	fallback string,
) string {
	if s := i.Argument(number); s != "" {
		return s
	}

	return fallback
}
