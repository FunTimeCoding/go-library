package concern

func New(
	key string,
	text string,
	line int,
	lineText string,
) *Concern {
	return &Concern{Key: key, Text: text, Line: line, LineText: lineText}
}
