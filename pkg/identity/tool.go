package identity

import "github.com/funtimecoding/go-library/pkg/identity/paragraph"

type Tool struct {
	name         string
	description  string
	usage        string
	instructions string
	paragraphs   []*paragraph.Paragraph
}
