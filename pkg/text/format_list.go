package text

import "github.com/funtimecoding/go-library/pkg/text/multi_line"

func FormatList(
	title string,
	elements []string,
) string {
	l := multi_line.New()
	l.Add("%s:", title)

	for _, element := range elements {
		l.Add("%s", element)
	}

	return l.Format()
}
