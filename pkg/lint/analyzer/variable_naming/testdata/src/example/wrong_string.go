package example

func WrongString() {
	x := "hello" // want `variable x of type string should be named s`
	_ = x
}
