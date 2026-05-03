package example

type Container struct {
	Nested struct { // want `anonymous struct; extract a named type`
		Value int
	}
}
