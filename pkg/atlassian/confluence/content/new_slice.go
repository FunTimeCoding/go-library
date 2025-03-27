package content

import kaos "github.com/essentialkaos/go-confluence/v6"

func NewSlice(
	v []*kaos.Content,
	host string,
) []*Content {
	var result []*Content

	for _, c := range v {
		result = append(result, New(c, host))
	}

	return result
}
