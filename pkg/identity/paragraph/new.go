package paragraph

func New(
	key string,
	text string,
) *Paragraph {
	return &Paragraph{Key: key, Text: text}
}
