package identity

import "github.com/funtimecoding/go-library/pkg/identity/paragraph"

func (t *Tool) WithInstructions(
	v string,
	paragraphs ...*paragraph.Paragraph,
) *Tool {
	t.instructions = v
	t.paragraphs = paragraphs

	return t
}
