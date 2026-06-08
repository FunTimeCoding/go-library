package report

type Section struct {
	title         string
	indent        int
	maximumLength int
	renderables   []renderable
}
