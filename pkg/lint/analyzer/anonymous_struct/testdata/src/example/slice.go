package example

var patterns = []struct { // want `anonymous struct; extract a named type`
	keyword string
	score   int
}{
	{"test", 1},
}
