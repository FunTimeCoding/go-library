package example

func LiteralAssign() string {
	s := "hello"
	s += "world" // want `use fmt\.Sprintf instead of string concatenation`

	return s
}
