package example

func OkCorrect(m map[string]int) {
	v, okay := m["key"] // want `variable v of type int should be named i`
	_ = v
	_ = okay
}
