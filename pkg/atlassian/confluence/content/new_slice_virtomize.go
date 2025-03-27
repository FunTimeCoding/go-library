package content

import virtomize "github.com/virtomize/confluence-go-api"

func NewSliceVirtomize(v []virtomize.Results) []*Content {
	var result []*Content

	for _, c := range v {
		result = append(result, NewVirtomize(&c))
	}

	return result
}
