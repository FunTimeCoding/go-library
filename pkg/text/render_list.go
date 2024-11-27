package text

import "github.com/funtimecoding/go-library/pkg/text/multi_line"

func RenderList(
	title string,
	elements []string,
) string {
	l := multi_line.New()
	l.Format("%s:", title)

	for _, element := range elements {
		l.Format("%s", element)
	}

	return l.Render()
}
