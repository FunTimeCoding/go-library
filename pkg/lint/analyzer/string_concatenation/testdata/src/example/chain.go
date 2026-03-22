package example

func Chain(a, b, c string) string {
	return a + b + c // want `use fmt\.Sprintf instead of string concatenation`
}
