package concern

func NewLine(
	key string,
	text string,
	path string,
	line int,
	lineText string,
	fixed bool,
) *Concern {
	return &Concern{
		Key:      key,
		Text:     text,
		Path:     path,
		Type:     Line,
		Line:     line,
		LineText: lineText,
		Fixed:    fixed,
	}
}
