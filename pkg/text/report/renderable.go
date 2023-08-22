package report

type renderable interface {
	Render() string
	Length() int
	indentDeeper()
}
