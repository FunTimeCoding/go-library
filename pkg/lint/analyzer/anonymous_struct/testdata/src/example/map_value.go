package example

type Registry struct {
	Entries map[string]struct { // want `anonymous struct; extract a named type`
		Comment string
	}
}
