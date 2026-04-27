package example

func ParamNonPrimitive(myMap map[string]int) { // want `variable myMap of type map\[string\]int should be named m`
	_ = myMap
}
