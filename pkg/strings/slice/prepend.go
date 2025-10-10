package slice

func Prepend(
	v []string,
	e string,
) []string {
	return append([]string{e}, v...)
}
