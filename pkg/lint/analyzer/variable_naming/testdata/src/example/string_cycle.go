package example

func StringCycle() {
	s := "first"
	p := "second" // want `variable p of type string should be named t`
	_ = s
	_ = p
}
