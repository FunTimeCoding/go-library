package content

import kaos "github.com/essentialkaos/go-confluence/v6"

func New(
	v *kaos.Content,
	host string,
) *Content {
	return &Content{
		Identifier: v.ID,
		Space:      parseSpace(v),
		Title:      v.Title,
		Link:       parseLink(v, host),
		TinyLink:   parseTinyLink(v, host),
		Labels:     parseLabels(v),
		Raw:        v,
	}
}
