package example_list

type Model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}
