package report

import "fmt"

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
