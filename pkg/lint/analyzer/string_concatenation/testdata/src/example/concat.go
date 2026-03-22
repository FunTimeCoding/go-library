package example

func Concat(a, b string) string {
	return a + b // want `use fmt\.Sprintf instead of string concatenation`
}
