package errors

func LogFlush(c Flusher) {
	LogOnError(c.Flush())
}
