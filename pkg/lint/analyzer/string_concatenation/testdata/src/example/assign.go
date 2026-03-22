package example

func Assign(a, b string) string {
	a += b // want `use fmt\.Sprintf instead of string concatenation`

	return a
}
