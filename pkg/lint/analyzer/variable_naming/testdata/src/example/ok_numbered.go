package example

func OkNumbered(m map[string]int) {
	v, ok := m["a"]  // want `variable v of type int should be named i` `variable ok should be named okay`
	w, ok2 := m["b"] // want `variable w of type int should be named a` `variable ok2 should be named okay1`
	_ = v
	_ = w
	_ = ok
	_ = ok2
}
