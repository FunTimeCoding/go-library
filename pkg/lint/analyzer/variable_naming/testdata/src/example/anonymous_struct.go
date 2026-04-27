package example

func AnonymousStruct() {
	x := struct{ Name string }{} // want `variable x of type struct\{Name string\} should be named s`
	_ = x
}
