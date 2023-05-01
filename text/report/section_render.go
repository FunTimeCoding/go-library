package report

import "fmt"

func (s *Section) Render() string {
	if len(s.renderables) == 0 {
		return ""
	}

	result := spaces(s.indent) + s.title

	for _, element := range s.renderables {
		result += fmt.Sprintf("\n%s", element.Render())
	}

	return result
}
