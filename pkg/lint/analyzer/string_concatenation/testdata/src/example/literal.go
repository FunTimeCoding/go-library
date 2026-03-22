package example

func Literal() string {
	return "hello" + "world" // want `use fmt\.Sprintf instead of string concatenation`
}
