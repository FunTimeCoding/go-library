package concern

func New(
	key string,
	text string,
	path string,
	line int,
	lineText string,
	fixed bool,
) *Concern {
	return NewLine(key, text, path, line, lineText, fixed)
}
