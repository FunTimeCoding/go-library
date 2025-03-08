package concern

func New(
	key string,
	text string,
	path string,
	line int,
	lineText string,
) *Concern {
	return &Concern{
		Key:      key,
		Text:     text,
		Path:     path,
		Line:     line,
		LineText: lineText,
	}
}
