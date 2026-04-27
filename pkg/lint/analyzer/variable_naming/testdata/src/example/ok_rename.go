package example

func OkRename(m map[string]int) {
	v, found := m["key"] // want `variable v of type int should be named i` `variable found should be named okay`
	_ = v
	_ = found
}
