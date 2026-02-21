package logger

func (l *Logger) Structured(
	text string,
	a ...any,
) {
	l.structured.Info(text, a...)
}
