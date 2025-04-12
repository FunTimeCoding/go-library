package toast

func New(
	identifier int,
	text string,
) *Toast {
	return &Toast{Identifier: identifier, Text: text}
}
