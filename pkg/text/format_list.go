package text

import "github.com/funtimecoding/go-library/pkg/text/multi_line"

func FormatList(
	title string,
	elements []string,
) string {
	result := multi_line.New()
	result.Add("%s:", title)

	for _, element := range elements {
		result.Add(element)
	}

	return result.Format()
}
