package log

func (l *Logger) Write(p []byte) (int, error) {
	l.writes <- p
	<-l.done

	return len(p), nil
}
