package logger

func (l *Logger) Plain(
	format string,
	a ...any,
) {
	l.plain.Printf(format, a...)
}
