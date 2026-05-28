package log

func (l *Logger) History() []string {
	mutex.Lock()
	defer mutex.Unlock()
	result := make([]string, len(l.history))
	copy(result, l.history)

	return result
}
