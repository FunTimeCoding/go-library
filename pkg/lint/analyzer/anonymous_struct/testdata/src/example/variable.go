package example

func Variable() any {
	var result struct { // want `anonymous struct; extract a named type`
		Name string
	}

	return result
}
