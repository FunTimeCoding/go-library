package content

import virtomize "github.com/virtomize/confluence-go-api"

func NewVirtomize(v *virtomize.Results) *Content {
	return &Content{
		Title:        v.Title,
		Link:         v.URL,
		RawVirtomize: v,
	}
}
