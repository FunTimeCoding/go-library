package report

import "fmt"

type Section struct {
	title         string
	indent        int
	maximumLength int
	renderables   []renderable
}

func (s *Section) indentDeeper() {
	s.indent += 1

	for _, e := range s.renderables {
		e.indentDeeper()
	}
}

func (s *Section) Length() int {
	if len(s.renderables) == 0 {
		return len(s.title)
	}

	return len(s.Render())
}

func (s *Section) Render() string {
	if len(s.renderables) == 0 {
		return ""
	}

	result := fmt.Sprintf("%s%s", spaces(s.indent), s.title)

	for _, e := range s.renderables {
		result = fmt.Sprintf("%s\n%s", result, e.Render())
	}

	return result
}
