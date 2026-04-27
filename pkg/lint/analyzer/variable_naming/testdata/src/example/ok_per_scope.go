package example

func OkPerScope(m map[string]int) {
	if v, ok := m["a"]; ok { // want `variable v of type int should be named i` `variable ok should be named okay`
		_ = v
	}

	if w, ok := m["b"]; ok { // want `variable w of type int should be named a` `variable ok should be named okay`
		_ = w
	}
}
