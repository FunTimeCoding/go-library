package example

func ParamPrimitiveWrong(x string) { // want `variable x of type string should be named s`
	_ = x
}
